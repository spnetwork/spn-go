package horizon

import (
	"github.com/spn/go/services/horizon/internal/db2"
	"github.com/spn/go/services/horizon/internal/db2/history"
	"github.com/spn/go/services/horizon/internal/ledger"
	"github.com/spn/go/services/horizon/internal/resourceadapter"
	"github.com/spn/go/protocols/horizon"
	"github.com/spn/go/support/render/hal"
	"github.com/spn/go/services/horizon/internal/render/problem"
	"github.com/spn/go/services/horizon/internal/render/sse"
)

// This file contains the actions:
//
// LedgerIndexAction: pages of ledgers
// LedgerShowAction: single ledger by sequence

// LedgerIndexAction renders a page of ledger resources, identified by
// a normal page query.
type LedgerIndexAction struct {
	Action
	PagingParams db2.PageQuery
	Records      []history.Ledger
	Page         hal.Page
}

// JSON is a method for actions.JSON
func (action *LedgerIndexAction) JSON() {
	action.Do(
		action.EnsureHistoryFreshness,
		action.loadParams,
		action.ValidateCursorWithinHistory,
		action.loadRecords,
		action.loadPage,
		func() { hal.Render(action.W, action.Page) },
	)
}

// SSE is a method for actions.SSE
func (action *LedgerIndexAction) SSE(stream sse.Stream) {
	action.Setup(
		action.EnsureHistoryFreshness,
		action.loadParams,
		action.ValidateCursorWithinHistory,
	)
	action.Do(
		action.loadRecords,
		func() {
			stream.SetLimit(int(action.PagingParams.Limit))
			records := action.Records[stream.SentCount():]

			for _, record := range records {
				var res horizon.Ledger
				resourceadapter.PopulateLedger(action.R.Context(), &res, record)
				stream.Send(sse.Event{ID: res.PagingToken(), Data: res})
			}
		},
	)
}

func (action *LedgerIndexAction) loadParams() {
	action.ValidateCursorAsDefault()
	action.PagingParams = action.GetPageQuery()
}

func (action *LedgerIndexAction) loadRecords() {
	action.Err = action.HistoryQ().Ledgers().
		Page(action.PagingParams).
		Select(&action.Records)
}

func (action *LedgerIndexAction) loadPage() {
	for _, record := range action.Records {
		var res horizon.Ledger
		resourceadapter.PopulateLedger(action.R.Context(), &res, record)
		action.Page.Add(res)
	}

	action.Page.FullURL = action.FullURL()
	action.Page.Limit = action.PagingParams.Limit
	action.Page.Cursor = action.PagingParams.Cursor
	action.Page.Order = action.PagingParams.Order
	action.Page.PopulateLinks()
}

// LedgerShowAction renders a ledger found by its sequence number.
type LedgerShowAction struct {
	Action
	Sequence int32
	Record   history.Ledger
}

// JSON is a method for actions.JSON
func (action *LedgerShowAction) JSON() {
	action.Do(
		action.EnsureHistoryFreshness,
		action.loadParams,
		action.verifyWithinHistory,
		action.loadRecord,
		func() {
			var res horizon.Ledger
			resourceadapter.PopulateLedger(action.R.Context(), &res, action.Record)
			hal.Render(action.W, res)
		},
	)
}

func (action *LedgerShowAction) loadParams() {
	action.Sequence = action.GetInt32("ledger_id")
}

func (action *LedgerShowAction) loadRecord() {
	action.Err = action.HistoryQ().
		LedgerBySequence(&action.Record, action.Sequence)
}

func (action *LedgerShowAction) verifyWithinHistory() {
	if action.Sequence < ledger.CurrentState().HistoryElder {
		action.Err = &problem.BeforeHistory
	}
}
