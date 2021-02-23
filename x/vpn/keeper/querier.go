package keeper

import (
	"context"

	"github.com/sentinel-official/hub/x/deposit"
	"github.com/sentinel-official/hub/x/node"
	"github.com/sentinel-official/hub/x/plan"
	"github.com/sentinel-official/hub/x/provider"
	"github.com/sentinel-official/hub/x/session"
	"github.com/sentinel-official/hub/x/subscription"
)

type Querier struct {
	Deposit      deposit.Querier
	Provider     provider.Querier
	Node         node.Querier
	Plan         plan.Querier
	Subscription subscription.Querier
	Session      session.Querier
}

func (q *Querier) QueryDeposit(c context.Context, req *deposit.QueryDepositRequest) (*deposit.QueryDepositResponse, error) {
	return q.Deposit.QueryDeposit(c, req)
}

func (q *Querier) QueryDeposits(c context.Context, req *deposit.QueryDepositsRequest) (*deposit.QueryDepositsResponse, error) {
	return q.Deposit.QueryDeposits(c, req)
}

func (q *Querier) QueryProvider(c context.Context, req *provider.QueryProviderRequest) (*provider.QueryProviderResponse, error) {
	return q.Provider.QueryProvider(c, req)
}

func (q *Querier) QueryProviders(c context.Context, req *provider.QueryProvidersRequest) (*provider.QueryProvidersResponse, error) {
	return q.Provider.QueryProviders(c, req)
}

func (q *Querier) QueryNode(c context.Context, req *node.QueryNodeRequest) (*node.QueryNodeResponse, error) {
	return q.Node.QueryNode(c, req)
}

func (q *Querier) QueryNodes(c context.Context, req *node.QueryNodesRequest) (res *node.QueryNodesResponse, err error) {
	return q.Node.QueryNodes(c, req)
}

func (q *Querier) QueryNodesForProvider(c context.Context, req *node.QueryNodesForProviderRequest) (res *node.QueryNodesForProviderResponse, err error) {
	return q.Node.QueryNodesForProvider(c, req)
}

func (q *Querier) QueryPlan(c context.Context, req *plan.QueryPlanRequest) (*plan.QueryPlanResponse, error) {
	return q.Plan.QueryPlan(c, req)
}

func (q *Querier) QueryPlans(c context.Context, req *plan.QueryPlansRequest) (res *plan.QueryPlansResponse, err error) {
	return q.Plan.QueryPlans(c, req)
}

func (q *Querier) QueryPlansForProvider(c context.Context, req *plan.QueryPlansForProviderRequest) (res *plan.QueryPlansForProviderResponse, err error) {
	return q.Plan.QueryPlansForProvider(c, req)
}

func (q *Querier) QueryNodesForPlan(c context.Context, req *plan.QueryNodesForPlanRequest) (*plan.QueryNodesForPlanResponse, error) {
	return q.Plan.QueryNodesForPlan(c, req)
}

func (q *Querier) QuerySubscription(c context.Context, req *subscription.QuerySubscriptionRequest) (*subscription.QuerySubscriptionResponse, error) {
	return q.Subscription.QuerySubscription(c, req)
}

func (q *Querier) QuerySubscriptions(c context.Context, req *subscription.QuerySubscriptionsRequest) (*subscription.QuerySubscriptionsResponse, error) {
	return q.Subscription.QuerySubscriptions(c, req)
}

func (q *Querier) QuerySubscriptionsForNode(c context.Context, req *subscription.QuerySubscriptionsForNodeRequest) (*subscription.QuerySubscriptionsForNodeResponse, error) {
	return q.Subscription.QuerySubscriptionsForNode(c, req)
}

func (q *Querier) QuerySubscriptionsForPlan(c context.Context, req *subscription.QuerySubscriptionsForPlanRequest) (*subscription.QuerySubscriptionsForPlanResponse, error) {
	return q.Subscription.QuerySubscriptionsForPlan(c, req)
}

func (q *Querier) QuerySubscriptionsForAddress(c context.Context, req *subscription.QuerySubscriptionsForAddressRequest) (*subscription.QuerySubscriptionsForAddressResponse, error) {
	return q.Subscription.QuerySubscriptionsForAddress(c, req)
}

func (q *Querier) QueryQuota(c context.Context, req *subscription.QueryQuotaRequest) (*subscription.QueryQuotaResponse, error) {
	return q.Subscription.QueryQuota(c, req)
}

func (q *Querier) QueryQuotas(c context.Context, req *subscription.QueryQuotasRequest) (*subscription.QueryQuotasResponse, error) {
	return q.Subscription.QueryQuotas(c, req)
}

func (q *Querier) QuerySession(c context.Context, req *session.QuerySessionRequest) (*session.QuerySessionResponse, error) {
	return q.Session.QuerySession(c, req)
}

func (q *Querier) QuerySessions(c context.Context, req *session.QuerySessionsRequest) (*session.QuerySessionsResponse, error) {
	return q.Session.QuerySessions(c, req)
}

func (q *Querier) QuerySessionsForSubscription(c context.Context, req *session.QuerySessionsForSubscriptionRequest) (*session.QuerySessionsForSubscriptionResponse, error) {
	return q.Session.QuerySessionsForSubscription(c, req)
}

func (q *Querier) QuerySessionsForNode(c context.Context, req *session.QuerySessionsForNodeRequest) (*session.QuerySessionsForNodeResponse, error) {
	return q.Session.QuerySessionsForNode(c, req)
}

func (q *Querier) QuerySessionsForAddress(c context.Context, req *session.QuerySessionsForAddressRequest) (*session.QuerySessionsForAddressResponse, error) {
	return q.Session.QuerySessionsForAddress(c, req)
}

func (q *Querier) QueryOngoingSession(c context.Context, req *session.QueryOngoingSessionRequest) (*session.QueryOngoingSessionResponse, error) {
	return q.Session.QueryOngoingSession(c, req)
}
