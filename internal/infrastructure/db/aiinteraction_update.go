// Code generated by ent, DO NOT EDIT.

package db

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/alfariiizi/vandor/internal/infrastructure/db/aiinteraction"
	"github.com/alfariiizi/vandor/internal/infrastructure/db/predicate"
)

// AIInteractionUpdate is the builder for updating AIInteraction entities.
type AIInteractionUpdate struct {
	config
	hooks    []Hook
	mutation *AIInteractionMutation
}

// Where appends a list predicates to the AIInteractionUpdate builder.
func (aiu *AIInteractionUpdate) Where(ps ...predicate.AIInteraction) *AIInteractionUpdate {
	aiu.mutation.Where(ps...)
	return aiu
}

// SetClinicID sets the "clinic_id" field.
func (aiu *AIInteractionUpdate) SetClinicID(s string) *AIInteractionUpdate {
	aiu.mutation.SetClinicID(s)
	return aiu
}

// SetNillableClinicID sets the "clinic_id" field if the given value is not nil.
func (aiu *AIInteractionUpdate) SetNillableClinicID(s *string) *AIInteractionUpdate {
	if s != nil {
		aiu.SetClinicID(*s)
	}
	return aiu
}

// SetPatientWhatsapp sets the "patient_whatsapp" field.
func (aiu *AIInteractionUpdate) SetPatientWhatsapp(s string) *AIInteractionUpdate {
	aiu.mutation.SetPatientWhatsapp(s)
	return aiu
}

// SetNillablePatientWhatsapp sets the "patient_whatsapp" field if the given value is not nil.
func (aiu *AIInteractionUpdate) SetNillablePatientWhatsapp(s *string) *AIInteractionUpdate {
	if s != nil {
		aiu.SetPatientWhatsapp(*s)
	}
	return aiu
}

// ClearPatientWhatsapp clears the value of the "patient_whatsapp" field.
func (aiu *AIInteractionUpdate) ClearPatientWhatsapp() *AIInteractionUpdate {
	aiu.mutation.ClearPatientWhatsapp()
	return aiu
}

// SetInteractionType sets the "interaction_type" field.
func (aiu *AIInteractionUpdate) SetInteractionType(at aiinteraction.InteractionType) *AIInteractionUpdate {
	aiu.mutation.SetInteractionType(at)
	return aiu
}

// SetNillableInteractionType sets the "interaction_type" field if the given value is not nil.
func (aiu *AIInteractionUpdate) SetNillableInteractionType(at *aiinteraction.InteractionType) *AIInteractionUpdate {
	if at != nil {
		aiu.SetInteractionType(*at)
	}
	return aiu
}

// SetRequestPayload sets the "request_payload" field.
func (aiu *AIInteractionUpdate) SetRequestPayload(m map[string]interface{}) *AIInteractionUpdate {
	aiu.mutation.SetRequestPayload(m)
	return aiu
}

// SetResponsePayload sets the "response_payload" field.
func (aiu *AIInteractionUpdate) SetResponsePayload(m map[string]interface{}) *AIInteractionUpdate {
	aiu.mutation.SetResponsePayload(m)
	return aiu
}

// SetAiModel sets the "ai_model" field.
func (aiu *AIInteractionUpdate) SetAiModel(s string) *AIInteractionUpdate {
	aiu.mutation.SetAiModel(s)
	return aiu
}

// SetNillableAiModel sets the "ai_model" field if the given value is not nil.
func (aiu *AIInteractionUpdate) SetNillableAiModel(s *string) *AIInteractionUpdate {
	if s != nil {
		aiu.SetAiModel(*s)
	}
	return aiu
}

// ClearAiModel clears the value of the "ai_model" field.
func (aiu *AIInteractionUpdate) ClearAiModel() *AIInteractionUpdate {
	aiu.mutation.ClearAiModel()
	return aiu
}

// SetResponseTimeMs sets the "response_time_ms" field.
func (aiu *AIInteractionUpdate) SetResponseTimeMs(i int) *AIInteractionUpdate {
	aiu.mutation.ResetResponseTimeMs()
	aiu.mutation.SetResponseTimeMs(i)
	return aiu
}

// SetNillableResponseTimeMs sets the "response_time_ms" field if the given value is not nil.
func (aiu *AIInteractionUpdate) SetNillableResponseTimeMs(i *int) *AIInteractionUpdate {
	if i != nil {
		aiu.SetResponseTimeMs(*i)
	}
	return aiu
}

// AddResponseTimeMs adds i to the "response_time_ms" field.
func (aiu *AIInteractionUpdate) AddResponseTimeMs(i int) *AIInteractionUpdate {
	aiu.mutation.AddResponseTimeMs(i)
	return aiu
}

// ClearResponseTimeMs clears the value of the "response_time_ms" field.
func (aiu *AIInteractionUpdate) ClearResponseTimeMs() *AIInteractionUpdate {
	aiu.mutation.ClearResponseTimeMs()
	return aiu
}

// SetStatus sets the "status" field.
func (aiu *AIInteractionUpdate) SetStatus(a aiinteraction.Status) *AIInteractionUpdate {
	aiu.mutation.SetStatus(a)
	return aiu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (aiu *AIInteractionUpdate) SetNillableStatus(a *aiinteraction.Status) *AIInteractionUpdate {
	if a != nil {
		aiu.SetStatus(*a)
	}
	return aiu
}

// SetErrorMessage sets the "error_message" field.
func (aiu *AIInteractionUpdate) SetErrorMessage(s string) *AIInteractionUpdate {
	aiu.mutation.SetErrorMessage(s)
	return aiu
}

// SetNillableErrorMessage sets the "error_message" field if the given value is not nil.
func (aiu *AIInteractionUpdate) SetNillableErrorMessage(s *string) *AIInteractionUpdate {
	if s != nil {
		aiu.SetErrorMessage(*s)
	}
	return aiu
}

// ClearErrorMessage clears the value of the "error_message" field.
func (aiu *AIInteractionUpdate) ClearErrorMessage() *AIInteractionUpdate {
	aiu.mutation.ClearErrorMessage()
	return aiu
}

// SetCreatedAt sets the "created_at" field.
func (aiu *AIInteractionUpdate) SetCreatedAt(t time.Time) *AIInteractionUpdate {
	aiu.mutation.SetCreatedAt(t)
	return aiu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (aiu *AIInteractionUpdate) SetNillableCreatedAt(t *time.Time) *AIInteractionUpdate {
	if t != nil {
		aiu.SetCreatedAt(*t)
	}
	return aiu
}

// Mutation returns the AIInteractionMutation object of the builder.
func (aiu *AIInteractionUpdate) Mutation() *AIInteractionMutation {
	return aiu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (aiu *AIInteractionUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, aiu.sqlSave, aiu.mutation, aiu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (aiu *AIInteractionUpdate) SaveX(ctx context.Context) int {
	affected, err := aiu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (aiu *AIInteractionUpdate) Exec(ctx context.Context) error {
	_, err := aiu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (aiu *AIInteractionUpdate) ExecX(ctx context.Context) {
	if err := aiu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (aiu *AIInteractionUpdate) check() error {
	if v, ok := aiu.mutation.InteractionType(); ok {
		if err := aiinteraction.InteractionTypeValidator(v); err != nil {
			return &ValidationError{Name: "interaction_type", err: fmt.Errorf(`db: validator failed for field "AIInteraction.interaction_type": %w`, err)}
		}
	}
	if v, ok := aiu.mutation.Status(); ok {
		if err := aiinteraction.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`db: validator failed for field "AIInteraction.status": %w`, err)}
		}
	}
	return nil
}

func (aiu *AIInteractionUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := aiu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(aiinteraction.Table, aiinteraction.Columns, sqlgraph.NewFieldSpec(aiinteraction.FieldID, field.TypeUUID))
	if ps := aiu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := aiu.mutation.ClinicID(); ok {
		_spec.SetField(aiinteraction.FieldClinicID, field.TypeString, value)
	}
	if value, ok := aiu.mutation.PatientWhatsapp(); ok {
		_spec.SetField(aiinteraction.FieldPatientWhatsapp, field.TypeString, value)
	}
	if aiu.mutation.PatientWhatsappCleared() {
		_spec.ClearField(aiinteraction.FieldPatientWhatsapp, field.TypeString)
	}
	if value, ok := aiu.mutation.InteractionType(); ok {
		_spec.SetField(aiinteraction.FieldInteractionType, field.TypeEnum, value)
	}
	if value, ok := aiu.mutation.RequestPayload(); ok {
		_spec.SetField(aiinteraction.FieldRequestPayload, field.TypeJSON, value)
	}
	if value, ok := aiu.mutation.ResponsePayload(); ok {
		_spec.SetField(aiinteraction.FieldResponsePayload, field.TypeJSON, value)
	}
	if value, ok := aiu.mutation.AiModel(); ok {
		_spec.SetField(aiinteraction.FieldAiModel, field.TypeString, value)
	}
	if aiu.mutation.AiModelCleared() {
		_spec.ClearField(aiinteraction.FieldAiModel, field.TypeString)
	}
	if value, ok := aiu.mutation.ResponseTimeMs(); ok {
		_spec.SetField(aiinteraction.FieldResponseTimeMs, field.TypeInt, value)
	}
	if value, ok := aiu.mutation.AddedResponseTimeMs(); ok {
		_spec.AddField(aiinteraction.FieldResponseTimeMs, field.TypeInt, value)
	}
	if aiu.mutation.ResponseTimeMsCleared() {
		_spec.ClearField(aiinteraction.FieldResponseTimeMs, field.TypeInt)
	}
	if value, ok := aiu.mutation.Status(); ok {
		_spec.SetField(aiinteraction.FieldStatus, field.TypeEnum, value)
	}
	if value, ok := aiu.mutation.ErrorMessage(); ok {
		_spec.SetField(aiinteraction.FieldErrorMessage, field.TypeString, value)
	}
	if aiu.mutation.ErrorMessageCleared() {
		_spec.ClearField(aiinteraction.FieldErrorMessage, field.TypeString)
	}
	if value, ok := aiu.mutation.CreatedAt(); ok {
		_spec.SetField(aiinteraction.FieldCreatedAt, field.TypeTime, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, aiu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{aiinteraction.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	aiu.mutation.done = true
	return n, nil
}

// AIInteractionUpdateOne is the builder for updating a single AIInteraction entity.
type AIInteractionUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *AIInteractionMutation
}

// SetClinicID sets the "clinic_id" field.
func (aiuo *AIInteractionUpdateOne) SetClinicID(s string) *AIInteractionUpdateOne {
	aiuo.mutation.SetClinicID(s)
	return aiuo
}

// SetNillableClinicID sets the "clinic_id" field if the given value is not nil.
func (aiuo *AIInteractionUpdateOne) SetNillableClinicID(s *string) *AIInteractionUpdateOne {
	if s != nil {
		aiuo.SetClinicID(*s)
	}
	return aiuo
}

// SetPatientWhatsapp sets the "patient_whatsapp" field.
func (aiuo *AIInteractionUpdateOne) SetPatientWhatsapp(s string) *AIInteractionUpdateOne {
	aiuo.mutation.SetPatientWhatsapp(s)
	return aiuo
}

// SetNillablePatientWhatsapp sets the "patient_whatsapp" field if the given value is not nil.
func (aiuo *AIInteractionUpdateOne) SetNillablePatientWhatsapp(s *string) *AIInteractionUpdateOne {
	if s != nil {
		aiuo.SetPatientWhatsapp(*s)
	}
	return aiuo
}

// ClearPatientWhatsapp clears the value of the "patient_whatsapp" field.
func (aiuo *AIInteractionUpdateOne) ClearPatientWhatsapp() *AIInteractionUpdateOne {
	aiuo.mutation.ClearPatientWhatsapp()
	return aiuo
}

// SetInteractionType sets the "interaction_type" field.
func (aiuo *AIInteractionUpdateOne) SetInteractionType(at aiinteraction.InteractionType) *AIInteractionUpdateOne {
	aiuo.mutation.SetInteractionType(at)
	return aiuo
}

// SetNillableInteractionType sets the "interaction_type" field if the given value is not nil.
func (aiuo *AIInteractionUpdateOne) SetNillableInteractionType(at *aiinteraction.InteractionType) *AIInteractionUpdateOne {
	if at != nil {
		aiuo.SetInteractionType(*at)
	}
	return aiuo
}

// SetRequestPayload sets the "request_payload" field.
func (aiuo *AIInteractionUpdateOne) SetRequestPayload(m map[string]interface{}) *AIInteractionUpdateOne {
	aiuo.mutation.SetRequestPayload(m)
	return aiuo
}

// SetResponsePayload sets the "response_payload" field.
func (aiuo *AIInteractionUpdateOne) SetResponsePayload(m map[string]interface{}) *AIInteractionUpdateOne {
	aiuo.mutation.SetResponsePayload(m)
	return aiuo
}

// SetAiModel sets the "ai_model" field.
func (aiuo *AIInteractionUpdateOne) SetAiModel(s string) *AIInteractionUpdateOne {
	aiuo.mutation.SetAiModel(s)
	return aiuo
}

// SetNillableAiModel sets the "ai_model" field if the given value is not nil.
func (aiuo *AIInteractionUpdateOne) SetNillableAiModel(s *string) *AIInteractionUpdateOne {
	if s != nil {
		aiuo.SetAiModel(*s)
	}
	return aiuo
}

// ClearAiModel clears the value of the "ai_model" field.
func (aiuo *AIInteractionUpdateOne) ClearAiModel() *AIInteractionUpdateOne {
	aiuo.mutation.ClearAiModel()
	return aiuo
}

// SetResponseTimeMs sets the "response_time_ms" field.
func (aiuo *AIInteractionUpdateOne) SetResponseTimeMs(i int) *AIInteractionUpdateOne {
	aiuo.mutation.ResetResponseTimeMs()
	aiuo.mutation.SetResponseTimeMs(i)
	return aiuo
}

// SetNillableResponseTimeMs sets the "response_time_ms" field if the given value is not nil.
func (aiuo *AIInteractionUpdateOne) SetNillableResponseTimeMs(i *int) *AIInteractionUpdateOne {
	if i != nil {
		aiuo.SetResponseTimeMs(*i)
	}
	return aiuo
}

// AddResponseTimeMs adds i to the "response_time_ms" field.
func (aiuo *AIInteractionUpdateOne) AddResponseTimeMs(i int) *AIInteractionUpdateOne {
	aiuo.mutation.AddResponseTimeMs(i)
	return aiuo
}

// ClearResponseTimeMs clears the value of the "response_time_ms" field.
func (aiuo *AIInteractionUpdateOne) ClearResponseTimeMs() *AIInteractionUpdateOne {
	aiuo.mutation.ClearResponseTimeMs()
	return aiuo
}

// SetStatus sets the "status" field.
func (aiuo *AIInteractionUpdateOne) SetStatus(a aiinteraction.Status) *AIInteractionUpdateOne {
	aiuo.mutation.SetStatus(a)
	return aiuo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (aiuo *AIInteractionUpdateOne) SetNillableStatus(a *aiinteraction.Status) *AIInteractionUpdateOne {
	if a != nil {
		aiuo.SetStatus(*a)
	}
	return aiuo
}

// SetErrorMessage sets the "error_message" field.
func (aiuo *AIInteractionUpdateOne) SetErrorMessage(s string) *AIInteractionUpdateOne {
	aiuo.mutation.SetErrorMessage(s)
	return aiuo
}

// SetNillableErrorMessage sets the "error_message" field if the given value is not nil.
func (aiuo *AIInteractionUpdateOne) SetNillableErrorMessage(s *string) *AIInteractionUpdateOne {
	if s != nil {
		aiuo.SetErrorMessage(*s)
	}
	return aiuo
}

// ClearErrorMessage clears the value of the "error_message" field.
func (aiuo *AIInteractionUpdateOne) ClearErrorMessage() *AIInteractionUpdateOne {
	aiuo.mutation.ClearErrorMessage()
	return aiuo
}

// SetCreatedAt sets the "created_at" field.
func (aiuo *AIInteractionUpdateOne) SetCreatedAt(t time.Time) *AIInteractionUpdateOne {
	aiuo.mutation.SetCreatedAt(t)
	return aiuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (aiuo *AIInteractionUpdateOne) SetNillableCreatedAt(t *time.Time) *AIInteractionUpdateOne {
	if t != nil {
		aiuo.SetCreatedAt(*t)
	}
	return aiuo
}

// Mutation returns the AIInteractionMutation object of the builder.
func (aiuo *AIInteractionUpdateOne) Mutation() *AIInteractionMutation {
	return aiuo.mutation
}

// Where appends a list predicates to the AIInteractionUpdate builder.
func (aiuo *AIInteractionUpdateOne) Where(ps ...predicate.AIInteraction) *AIInteractionUpdateOne {
	aiuo.mutation.Where(ps...)
	return aiuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (aiuo *AIInteractionUpdateOne) Select(field string, fields ...string) *AIInteractionUpdateOne {
	aiuo.fields = append([]string{field}, fields...)
	return aiuo
}

// Save executes the query and returns the updated AIInteraction entity.
func (aiuo *AIInteractionUpdateOne) Save(ctx context.Context) (*AIInteraction, error) {
	return withHooks(ctx, aiuo.sqlSave, aiuo.mutation, aiuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (aiuo *AIInteractionUpdateOne) SaveX(ctx context.Context) *AIInteraction {
	node, err := aiuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (aiuo *AIInteractionUpdateOne) Exec(ctx context.Context) error {
	_, err := aiuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (aiuo *AIInteractionUpdateOne) ExecX(ctx context.Context) {
	if err := aiuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (aiuo *AIInteractionUpdateOne) check() error {
	if v, ok := aiuo.mutation.InteractionType(); ok {
		if err := aiinteraction.InteractionTypeValidator(v); err != nil {
			return &ValidationError{Name: "interaction_type", err: fmt.Errorf(`db: validator failed for field "AIInteraction.interaction_type": %w`, err)}
		}
	}
	if v, ok := aiuo.mutation.Status(); ok {
		if err := aiinteraction.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`db: validator failed for field "AIInteraction.status": %w`, err)}
		}
	}
	return nil
}

func (aiuo *AIInteractionUpdateOne) sqlSave(ctx context.Context) (_node *AIInteraction, err error) {
	if err := aiuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(aiinteraction.Table, aiinteraction.Columns, sqlgraph.NewFieldSpec(aiinteraction.FieldID, field.TypeUUID))
	id, ok := aiuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`db: missing "AIInteraction.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := aiuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, aiinteraction.FieldID)
		for _, f := range fields {
			if !aiinteraction.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("db: invalid field %q for query", f)}
			}
			if f != aiinteraction.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := aiuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := aiuo.mutation.ClinicID(); ok {
		_spec.SetField(aiinteraction.FieldClinicID, field.TypeString, value)
	}
	if value, ok := aiuo.mutation.PatientWhatsapp(); ok {
		_spec.SetField(aiinteraction.FieldPatientWhatsapp, field.TypeString, value)
	}
	if aiuo.mutation.PatientWhatsappCleared() {
		_spec.ClearField(aiinteraction.FieldPatientWhatsapp, field.TypeString)
	}
	if value, ok := aiuo.mutation.InteractionType(); ok {
		_spec.SetField(aiinteraction.FieldInteractionType, field.TypeEnum, value)
	}
	if value, ok := aiuo.mutation.RequestPayload(); ok {
		_spec.SetField(aiinteraction.FieldRequestPayload, field.TypeJSON, value)
	}
	if value, ok := aiuo.mutation.ResponsePayload(); ok {
		_spec.SetField(aiinteraction.FieldResponsePayload, field.TypeJSON, value)
	}
	if value, ok := aiuo.mutation.AiModel(); ok {
		_spec.SetField(aiinteraction.FieldAiModel, field.TypeString, value)
	}
	if aiuo.mutation.AiModelCleared() {
		_spec.ClearField(aiinteraction.FieldAiModel, field.TypeString)
	}
	if value, ok := aiuo.mutation.ResponseTimeMs(); ok {
		_spec.SetField(aiinteraction.FieldResponseTimeMs, field.TypeInt, value)
	}
	if value, ok := aiuo.mutation.AddedResponseTimeMs(); ok {
		_spec.AddField(aiinteraction.FieldResponseTimeMs, field.TypeInt, value)
	}
	if aiuo.mutation.ResponseTimeMsCleared() {
		_spec.ClearField(aiinteraction.FieldResponseTimeMs, field.TypeInt)
	}
	if value, ok := aiuo.mutation.Status(); ok {
		_spec.SetField(aiinteraction.FieldStatus, field.TypeEnum, value)
	}
	if value, ok := aiuo.mutation.ErrorMessage(); ok {
		_spec.SetField(aiinteraction.FieldErrorMessage, field.TypeString, value)
	}
	if aiuo.mutation.ErrorMessageCleared() {
		_spec.ClearField(aiinteraction.FieldErrorMessage, field.TypeString)
	}
	if value, ok := aiuo.mutation.CreatedAt(); ok {
		_spec.SetField(aiinteraction.FieldCreatedAt, field.TypeTime, value)
	}
	_node = &AIInteraction{config: aiuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, aiuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{aiinteraction.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	aiuo.mutation.done = true
	return _node, nil
}
