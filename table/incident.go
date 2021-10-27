package table

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/madsaune/snowy/auth"
	"github.com/madsaune/snowy/client"
	"github.com/madsaune/snowy/fieldtype"
)

const (
	defaultEndpoint = "/api/now/table/incident"
)

type IncidentClient struct {
	BaseClient client.Client
}

func NewIncidentClient(authorizer auth.Authorizer) *IncidentClient {
	return &IncidentClient{
		BaseClient: client.NewClient(authorizer),
	}
}

func (i *IncidentClient) Get(ctx context.Context, id string) (*Incident, error) {
	u := fmt.Sprintf("%s/%s", defaultEndpoint, id)
	res, err := i.BaseClient.Get(ctx, u, nil)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var incident *IncidentResponse
	err = json.NewDecoder(res.Body).Decode(&incident)
	if err != nil {
		return nil, err
	}

	return &incident.Result, nil
}

type IncidentResponse struct {
	Result Incident `json:"result"`
}

type Incident struct {
	PromotedBy             string            `json:"promoted_by,omitempty"`
	Parent                 string            `json:"parent,omitempty"`
	CausedBy               string            `json:"caused_by,omitempty"`
	WatchList              string            `json:"watch_list,omitempty"`
	UponReject             string            `json:"upon_reject,omitempty"`
	SysUpdatedOn           fieldtype.SNTime  `json:"sys_updated_on,omitempty"`
	ApprovalHistory        string            `json:"approval_history,omitempty"`
	Skills                 string            `json:"skills,omitempty"`
	Number                 string            `json:"number,omitempty"`
	ProposedBy             string            `json:"proposed_by,omitempty"`
	LessonsLearned         string            `json:"lessons_learned,omitempty"`
	State                  string            `json:"state,omitempty"`
	SysCreatedBy           string            `json:"sys_created_by,omitempty"`
	Knowledge              string            `json:"knowledge,omitempty"`
	Order                  string            `json:"order,omitempty"`
	CmdbCI                 string            `json:"cmdb_ci,omitempty"`
	DeliveryPlan           string            `json:"delivery_plan,omitempty"`
	Impact                 string            `json:"impact,omitempty"`
	Active                 fieldtype.SNBool  `json:"active,omitempty"`
	WorkNotesList          string            `json:"work_notes_list,omitempty"`
	Priority               string            `json:"priority,omitempty"`
	SysDomainPath          string            `json:"sys_domain_path,omitempty"`
	BusinessDuration       string            `json:"business_duration,omitempty"`
	GroupList              string            `json:"group_list,omitempty"`
	ApprovalSet            string            `json:"approval_set,omitempty"`
	MajorIncidentState     string            `json:"major_incident_state,omitempty"`
	UniversalRequest       string            `json:"universal_request,omitempty"`
	ShortDescription       string            `json:"short_description,omitempty"`
	CorrelationDisplay     string            `json:"correlation_display,omitempty"`
	DeliveryTask           string            `json:"delivery_task,omitempty"`
	WorkStart              string            `json:"work_start,omitempty"`
	TriggerRule            string            `json:"trigger_rule,omitempty"`
	AdditionalAssigneeList string            `json:"additional_assignee_list,omitempty"`
	Notify                 string            `json:"notify,omitempty"`
	ServiceOffering        string            `json:"service_offering,omitempty"`
	SysClassName           string            `json:"sys_class_name,omitempty"`
	ClosedBy               string            `json:"closed_by,omitempty"`
	FollowUp               string            `json:"follow_up,omitempty"`
	ParentIncident         string            `json:"parent_incident,omitempty"`
	ReopenedBy             string            `json:"reopened_by,omitempty"`
	ReassignmentCount      fieldtype.SNInt   `json:"reassignment_count,omitempty"`
	AssignedTo             ExpandedParameter `json:"assigned_to,omitempty"`
	SlaDue                 string            `json:"sla_due,omitempty"`
	CommentsAndWorkNotes   string            `json:"comments_and_work_notes,omitempty"`
	Escalation             string            `json:"escalation,omitempty"`
	UponApproval           string            `json:"upon_approval,omitempty"`
	CorrelationID          string            `json:"correlation_id,omitempty"`
	Timeline               string            `json:"timeline,omitempty"`
	MadeSLA                string            `json:"made_sla,omitempty"`
	PromotedOn             string            `json:"promoted_on,omitempty"`
	ChildIncidents         string            `json:"child_incidents,omitempty"`
	HoldReason             string            `json:"hold_reason,omitempty"`
	TaskEffectiveNumber    string            `json:"task_effective_number,omitempty"`
	ResolvedBy             string            `json:"resolved_by,omitempty"`
	SysUpdatedBy           string            `json:"sys_updated_by,omitempty"`
	OpenedBy               ExpandedParameter `json:"opened_by,omitempty"`
	UserInput              string            `json:"user_input,omitempty"`
	SysCreatedOn           fieldtype.SNTime  `json:"sys_created_on,omitempty"`
	SysDomain              ExpandedParameter `json:"sys_domain,omitempty"`
	ProposedOn             string            `json:"proposed_on,omitempty"`
	ActionsTaken           string            `json:"actions_taken,omitempty"`
	TaskFor                ExpandedParameter `json:""task_for,omitempty"`
	RouteReason            string            `json:"route_reason,omitempty"`
	CalendarSTC            string            `json:"calendar_stc,omitempty"`
	ClosedAt               string            `json:"closed_at,omitempty"`
	BusinessService        string            `json:"business_service,omitempty"`
	BusinessImpact         string            `json:"business_impact,omitempty"`
	RFC                    string            `json:"rfc,omitempty"`
	TimeWorked             string            `json:"time_worked,omitempty"`
	ExpectedStart          string            `json:"expected_start,omitempty"`
	OpenedAt               fieldtype.SNTime  `json:"opened_at,omitempty"`
	WorkEnd                string            `json:"work_end,omitempty"`
	CallerID               ExpandedParameter `json:"caller_id,omitempty"`
	ReopenedTime           string            `json:"reopened_time,omitempty"`
	ResolvedAt             string            `json:"resolved_at,omitempty"`
	Subcategory            string            `json:"subcategory,omitempty"`
	WorkNotes              string            `json:"work_notes,omitempty"`
	CloseCode              string            `json:"close_code,omitempty"`
	AssignmentGroup        ExpandedParameter `json:"assignment_group,omitempty"`
	BusinessSTC            string            `json:"business_stc,omitempty"`
	Cause                  string            `json:"cause,omitempty"`
	Description            string            `json:"description,omitempty"`
	CalendarDuration       string            `json:"calendar_duration,omitempty"`
	CloseNotes             string            `json:"close_notes,omitempty"`
	SysID                  string            `json:"sys_id,omitempty"`
	ContactType            string            `json:"contact_type,omitempty"`
	IncidentState          string            `json:"incident_state,omitempty"`
	Urgency                string            `json:"urgency,omitempty"`
	ProblemID              string            `json:"problem_id,omitempty"`
	Company                ExpandedParameter `json:"company,omitempty"`
	Activitydue            string            `json:"activity_due,omitempty"`
	Severity               string            `json:"severity,omitempty"`
	Overview               string            `json:"overview,omitempty"`
	Comments               string            `json:"comments,omitempty"`
	Approval               string            `json:"approval,omitempty"`
	DueDate                string            `json:"due_date,omitempty"`
	SysModCount            string            `json:"sys_mod_count,omitempty"`
	ReopenCount            string            `json:"reopen_count,omitempty"`
	SysTags                string            `json:"sys_tags,omitempty"`
	Location               string            `json:"location,omitempty"`
	Category               string            `json:"category,omitempty"`
}

type ExpandedParameter struct {
	Link  string `json:"link,omitempty"`
	Value string `json:"value,omitempty"`
}
