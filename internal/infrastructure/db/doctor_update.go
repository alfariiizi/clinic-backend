// Code generated by ent, DO NOT EDIT.

package db

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/dialect/sql/sqljson"
	"entgo.io/ent/schema/field"
	"github.com/alfariiizi/vandor/internal/infrastructure/db/appointment"
	"github.com/alfariiizi/vandor/internal/infrastructure/db/clinic"
	"github.com/alfariiizi/vandor/internal/infrastructure/db/doctor"
	"github.com/alfariiizi/vandor/internal/infrastructure/db/doctorschedule"
	"github.com/alfariiizi/vandor/internal/infrastructure/db/predicate"
	"github.com/google/uuid"
)

// DoctorUpdate is the builder for updating Doctor entities.
type DoctorUpdate struct {
	config
	hooks    []Hook
	mutation *DoctorMutation
}

// Where appends a list predicates to the DoctorUpdate builder.
func (du *DoctorUpdate) Where(ps ...predicate.Doctor) *DoctorUpdate {
	du.mutation.Where(ps...)
	return du
}

// SetName sets the "name" field.
func (du *DoctorUpdate) SetName(s string) *DoctorUpdate {
	du.mutation.SetName(s)
	return du
}

// SetNillableName sets the "name" field if the given value is not nil.
func (du *DoctorUpdate) SetNillableName(s *string) *DoctorUpdate {
	if s != nil {
		du.SetName(*s)
	}
	return du
}

// SetSpecialization sets the "specialization" field.
func (du *DoctorUpdate) SetSpecialization(s string) *DoctorUpdate {
	du.mutation.SetSpecialization(s)
	return du
}

// SetNillableSpecialization sets the "specialization" field if the given value is not nil.
func (du *DoctorUpdate) SetNillableSpecialization(s *string) *DoctorUpdate {
	if s != nil {
		du.SetSpecialization(*s)
	}
	return du
}

// SetLicenseNumber sets the "license_number" field.
func (du *DoctorUpdate) SetLicenseNumber(s string) *DoctorUpdate {
	du.mutation.SetLicenseNumber(s)
	return du
}

// SetNillableLicenseNumber sets the "license_number" field if the given value is not nil.
func (du *DoctorUpdate) SetNillableLicenseNumber(s *string) *DoctorUpdate {
	if s != nil {
		du.SetLicenseNumber(*s)
	}
	return du
}

// ClearLicenseNumber clears the value of the "license_number" field.
func (du *DoctorUpdate) ClearLicenseNumber() *DoctorUpdate {
	du.mutation.ClearLicenseNumber()
	return du
}

// SetEmail sets the "email" field.
func (du *DoctorUpdate) SetEmail(s string) *DoctorUpdate {
	du.mutation.SetEmail(s)
	return du
}

// SetNillableEmail sets the "email" field if the given value is not nil.
func (du *DoctorUpdate) SetNillableEmail(s *string) *DoctorUpdate {
	if s != nil {
		du.SetEmail(*s)
	}
	return du
}

// ClearEmail clears the value of the "email" field.
func (du *DoctorUpdate) ClearEmail() *DoctorUpdate {
	du.mutation.ClearEmail()
	return du
}

// SetPhone sets the "phone" field.
func (du *DoctorUpdate) SetPhone(s string) *DoctorUpdate {
	du.mutation.SetPhone(s)
	return du
}

// SetNillablePhone sets the "phone" field if the given value is not nil.
func (du *DoctorUpdate) SetNillablePhone(s *string) *DoctorUpdate {
	if s != nil {
		du.SetPhone(*s)
	}
	return du
}

// ClearPhone clears the value of the "phone" field.
func (du *DoctorUpdate) ClearPhone() *DoctorUpdate {
	du.mutation.ClearPhone()
	return du
}

// SetBio sets the "bio" field.
func (du *DoctorUpdate) SetBio(s string) *DoctorUpdate {
	du.mutation.SetBio(s)
	return du
}

// SetNillableBio sets the "bio" field if the given value is not nil.
func (du *DoctorUpdate) SetNillableBio(s *string) *DoctorUpdate {
	if s != nil {
		du.SetBio(*s)
	}
	return du
}

// ClearBio clears the value of the "bio" field.
func (du *DoctorUpdate) ClearBio() *DoctorUpdate {
	du.mutation.ClearBio()
	return du
}

// SetQualifications sets the "qualifications" field.
func (du *DoctorUpdate) SetQualifications(s []string) *DoctorUpdate {
	du.mutation.SetQualifications(s)
	return du
}

// AppendQualifications appends s to the "qualifications" field.
func (du *DoctorUpdate) AppendQualifications(s []string) *DoctorUpdate {
	du.mutation.AppendQualifications(s)
	return du
}

// ClearQualifications clears the value of the "qualifications" field.
func (du *DoctorUpdate) ClearQualifications() *DoctorUpdate {
	du.mutation.ClearQualifications()
	return du
}

// SetAvailability sets the "availability" field.
func (du *DoctorUpdate) SetAvailability(m map[string]interface{}) *DoctorUpdate {
	du.mutation.SetAvailability(m)
	return du
}

// SetConsultationDuration sets the "consultation_duration" field.
func (du *DoctorUpdate) SetConsultationDuration(i int) *DoctorUpdate {
	du.mutation.ResetConsultationDuration()
	du.mutation.SetConsultationDuration(i)
	return du
}

// SetNillableConsultationDuration sets the "consultation_duration" field if the given value is not nil.
func (du *DoctorUpdate) SetNillableConsultationDuration(i *int) *DoctorUpdate {
	if i != nil {
		du.SetConsultationDuration(*i)
	}
	return du
}

// AddConsultationDuration adds i to the "consultation_duration" field.
func (du *DoctorUpdate) AddConsultationDuration(i int) *DoctorUpdate {
	du.mutation.AddConsultationDuration(i)
	return du
}

// SetConsultationFee sets the "consultation_fee" field.
func (du *DoctorUpdate) SetConsultationFee(f float64) *DoctorUpdate {
	du.mutation.ResetConsultationFee()
	du.mutation.SetConsultationFee(f)
	return du
}

// SetNillableConsultationFee sets the "consultation_fee" field if the given value is not nil.
func (du *DoctorUpdate) SetNillableConsultationFee(f *float64) *DoctorUpdate {
	if f != nil {
		du.SetConsultationFee(*f)
	}
	return du
}

// AddConsultationFee adds f to the "consultation_fee" field.
func (du *DoctorUpdate) AddConsultationFee(f float64) *DoctorUpdate {
	du.mutation.AddConsultationFee(f)
	return du
}

// SetActive sets the "active" field.
func (du *DoctorUpdate) SetActive(b bool) *DoctorUpdate {
	du.mutation.SetActive(b)
	return du
}

// SetNillableActive sets the "active" field if the given value is not nil.
func (du *DoctorUpdate) SetNillableActive(b *bool) *DoctorUpdate {
	if b != nil {
		du.SetActive(*b)
	}
	return du
}

// SetCreatedAt sets the "created_at" field.
func (du *DoctorUpdate) SetCreatedAt(t time.Time) *DoctorUpdate {
	du.mutation.SetCreatedAt(t)
	return du
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (du *DoctorUpdate) SetNillableCreatedAt(t *time.Time) *DoctorUpdate {
	if t != nil {
		du.SetCreatedAt(*t)
	}
	return du
}

// SetUpdatedAt sets the "updated_at" field.
func (du *DoctorUpdate) SetUpdatedAt(t time.Time) *DoctorUpdate {
	du.mutation.SetUpdatedAt(t)
	return du
}

// SetClinicID sets the "clinic" edge to the Clinic entity by ID.
func (du *DoctorUpdate) SetClinicID(id uuid.UUID) *DoctorUpdate {
	du.mutation.SetClinicID(id)
	return du
}

// SetNillableClinicID sets the "clinic" edge to the Clinic entity by ID if the given value is not nil.
func (du *DoctorUpdate) SetNillableClinicID(id *uuid.UUID) *DoctorUpdate {
	if id != nil {
		du = du.SetClinicID(*id)
	}
	return du
}

// SetClinic sets the "clinic" edge to the Clinic entity.
func (du *DoctorUpdate) SetClinic(c *Clinic) *DoctorUpdate {
	return du.SetClinicID(c.ID)
}

// AddAppointmentIDs adds the "appointments" edge to the Appointment entity by IDs.
func (du *DoctorUpdate) AddAppointmentIDs(ids ...uuid.UUID) *DoctorUpdate {
	du.mutation.AddAppointmentIDs(ids...)
	return du
}

// AddAppointments adds the "appointments" edges to the Appointment entity.
func (du *DoctorUpdate) AddAppointments(a ...*Appointment) *DoctorUpdate {
	ids := make([]uuid.UUID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return du.AddAppointmentIDs(ids...)
}

// AddScheduleIDs adds the "schedules" edge to the DoctorSchedule entity by IDs.
func (du *DoctorUpdate) AddScheduleIDs(ids ...uuid.UUID) *DoctorUpdate {
	du.mutation.AddScheduleIDs(ids...)
	return du
}

// AddSchedules adds the "schedules" edges to the DoctorSchedule entity.
func (du *DoctorUpdate) AddSchedules(d ...*DoctorSchedule) *DoctorUpdate {
	ids := make([]uuid.UUID, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return du.AddScheduleIDs(ids...)
}

// Mutation returns the DoctorMutation object of the builder.
func (du *DoctorUpdate) Mutation() *DoctorMutation {
	return du.mutation
}

// ClearClinic clears the "clinic" edge to the Clinic entity.
func (du *DoctorUpdate) ClearClinic() *DoctorUpdate {
	du.mutation.ClearClinic()
	return du
}

// ClearAppointments clears all "appointments" edges to the Appointment entity.
func (du *DoctorUpdate) ClearAppointments() *DoctorUpdate {
	du.mutation.ClearAppointments()
	return du
}

// RemoveAppointmentIDs removes the "appointments" edge to Appointment entities by IDs.
func (du *DoctorUpdate) RemoveAppointmentIDs(ids ...uuid.UUID) *DoctorUpdate {
	du.mutation.RemoveAppointmentIDs(ids...)
	return du
}

// RemoveAppointments removes "appointments" edges to Appointment entities.
func (du *DoctorUpdate) RemoveAppointments(a ...*Appointment) *DoctorUpdate {
	ids := make([]uuid.UUID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return du.RemoveAppointmentIDs(ids...)
}

// ClearSchedules clears all "schedules" edges to the DoctorSchedule entity.
func (du *DoctorUpdate) ClearSchedules() *DoctorUpdate {
	du.mutation.ClearSchedules()
	return du
}

// RemoveScheduleIDs removes the "schedules" edge to DoctorSchedule entities by IDs.
func (du *DoctorUpdate) RemoveScheduleIDs(ids ...uuid.UUID) *DoctorUpdate {
	du.mutation.RemoveScheduleIDs(ids...)
	return du
}

// RemoveSchedules removes "schedules" edges to DoctorSchedule entities.
func (du *DoctorUpdate) RemoveSchedules(d ...*DoctorSchedule) *DoctorUpdate {
	ids := make([]uuid.UUID, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return du.RemoveScheduleIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (du *DoctorUpdate) Save(ctx context.Context) (int, error) {
	du.defaults()
	return withHooks(ctx, du.sqlSave, du.mutation, du.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (du *DoctorUpdate) SaveX(ctx context.Context) int {
	affected, err := du.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (du *DoctorUpdate) Exec(ctx context.Context) error {
	_, err := du.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (du *DoctorUpdate) ExecX(ctx context.Context) {
	if err := du.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (du *DoctorUpdate) defaults() {
	if _, ok := du.mutation.UpdatedAt(); !ok {
		v := doctor.UpdateDefaultUpdatedAt()
		du.mutation.SetUpdatedAt(v)
	}
}

func (du *DoctorUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(doctor.Table, doctor.Columns, sqlgraph.NewFieldSpec(doctor.FieldID, field.TypeUUID))
	if ps := du.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := du.mutation.Name(); ok {
		_spec.SetField(doctor.FieldName, field.TypeString, value)
	}
	if value, ok := du.mutation.Specialization(); ok {
		_spec.SetField(doctor.FieldSpecialization, field.TypeString, value)
	}
	if value, ok := du.mutation.LicenseNumber(); ok {
		_spec.SetField(doctor.FieldLicenseNumber, field.TypeString, value)
	}
	if du.mutation.LicenseNumberCleared() {
		_spec.ClearField(doctor.FieldLicenseNumber, field.TypeString)
	}
	if value, ok := du.mutation.Email(); ok {
		_spec.SetField(doctor.FieldEmail, field.TypeString, value)
	}
	if du.mutation.EmailCleared() {
		_spec.ClearField(doctor.FieldEmail, field.TypeString)
	}
	if value, ok := du.mutation.Phone(); ok {
		_spec.SetField(doctor.FieldPhone, field.TypeString, value)
	}
	if du.mutation.PhoneCleared() {
		_spec.ClearField(doctor.FieldPhone, field.TypeString)
	}
	if value, ok := du.mutation.Bio(); ok {
		_spec.SetField(doctor.FieldBio, field.TypeString, value)
	}
	if du.mutation.BioCleared() {
		_spec.ClearField(doctor.FieldBio, field.TypeString)
	}
	if value, ok := du.mutation.Qualifications(); ok {
		_spec.SetField(doctor.FieldQualifications, field.TypeJSON, value)
	}
	if value, ok := du.mutation.AppendedQualifications(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, doctor.FieldQualifications, value)
		})
	}
	if du.mutation.QualificationsCleared() {
		_spec.ClearField(doctor.FieldQualifications, field.TypeJSON)
	}
	if value, ok := du.mutation.Availability(); ok {
		_spec.SetField(doctor.FieldAvailability, field.TypeJSON, value)
	}
	if value, ok := du.mutation.ConsultationDuration(); ok {
		_spec.SetField(doctor.FieldConsultationDuration, field.TypeInt, value)
	}
	if value, ok := du.mutation.AddedConsultationDuration(); ok {
		_spec.AddField(doctor.FieldConsultationDuration, field.TypeInt, value)
	}
	if value, ok := du.mutation.ConsultationFee(); ok {
		_spec.SetField(doctor.FieldConsultationFee, field.TypeFloat64, value)
	}
	if value, ok := du.mutation.AddedConsultationFee(); ok {
		_spec.AddField(doctor.FieldConsultationFee, field.TypeFloat64, value)
	}
	if value, ok := du.mutation.Active(); ok {
		_spec.SetField(doctor.FieldActive, field.TypeBool, value)
	}
	if value, ok := du.mutation.CreatedAt(); ok {
		_spec.SetField(doctor.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := du.mutation.UpdatedAt(); ok {
		_spec.SetField(doctor.FieldUpdatedAt, field.TypeTime, value)
	}
	if du.mutation.ClinicCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   doctor.ClinicTable,
			Columns: []string{doctor.ClinicColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(clinic.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := du.mutation.ClinicIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   doctor.ClinicTable,
			Columns: []string{doctor.ClinicColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(clinic.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if du.mutation.AppointmentsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   doctor.AppointmentsTable,
			Columns: []string{doctor.AppointmentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(appointment.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := du.mutation.RemovedAppointmentsIDs(); len(nodes) > 0 && !du.mutation.AppointmentsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   doctor.AppointmentsTable,
			Columns: []string{doctor.AppointmentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(appointment.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := du.mutation.AppointmentsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   doctor.AppointmentsTable,
			Columns: []string{doctor.AppointmentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(appointment.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if du.mutation.SchedulesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   doctor.SchedulesTable,
			Columns: []string{doctor.SchedulesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(doctorschedule.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := du.mutation.RemovedSchedulesIDs(); len(nodes) > 0 && !du.mutation.SchedulesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   doctor.SchedulesTable,
			Columns: []string{doctor.SchedulesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(doctorschedule.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := du.mutation.SchedulesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   doctor.SchedulesTable,
			Columns: []string{doctor.SchedulesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(doctorschedule.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, du.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{doctor.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	du.mutation.done = true
	return n, nil
}

// DoctorUpdateOne is the builder for updating a single Doctor entity.
type DoctorUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *DoctorMutation
}

// SetName sets the "name" field.
func (duo *DoctorUpdateOne) SetName(s string) *DoctorUpdateOne {
	duo.mutation.SetName(s)
	return duo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (duo *DoctorUpdateOne) SetNillableName(s *string) *DoctorUpdateOne {
	if s != nil {
		duo.SetName(*s)
	}
	return duo
}

// SetSpecialization sets the "specialization" field.
func (duo *DoctorUpdateOne) SetSpecialization(s string) *DoctorUpdateOne {
	duo.mutation.SetSpecialization(s)
	return duo
}

// SetNillableSpecialization sets the "specialization" field if the given value is not nil.
func (duo *DoctorUpdateOne) SetNillableSpecialization(s *string) *DoctorUpdateOne {
	if s != nil {
		duo.SetSpecialization(*s)
	}
	return duo
}

// SetLicenseNumber sets the "license_number" field.
func (duo *DoctorUpdateOne) SetLicenseNumber(s string) *DoctorUpdateOne {
	duo.mutation.SetLicenseNumber(s)
	return duo
}

// SetNillableLicenseNumber sets the "license_number" field if the given value is not nil.
func (duo *DoctorUpdateOne) SetNillableLicenseNumber(s *string) *DoctorUpdateOne {
	if s != nil {
		duo.SetLicenseNumber(*s)
	}
	return duo
}

// ClearLicenseNumber clears the value of the "license_number" field.
func (duo *DoctorUpdateOne) ClearLicenseNumber() *DoctorUpdateOne {
	duo.mutation.ClearLicenseNumber()
	return duo
}

// SetEmail sets the "email" field.
func (duo *DoctorUpdateOne) SetEmail(s string) *DoctorUpdateOne {
	duo.mutation.SetEmail(s)
	return duo
}

// SetNillableEmail sets the "email" field if the given value is not nil.
func (duo *DoctorUpdateOne) SetNillableEmail(s *string) *DoctorUpdateOne {
	if s != nil {
		duo.SetEmail(*s)
	}
	return duo
}

// ClearEmail clears the value of the "email" field.
func (duo *DoctorUpdateOne) ClearEmail() *DoctorUpdateOne {
	duo.mutation.ClearEmail()
	return duo
}

// SetPhone sets the "phone" field.
func (duo *DoctorUpdateOne) SetPhone(s string) *DoctorUpdateOne {
	duo.mutation.SetPhone(s)
	return duo
}

// SetNillablePhone sets the "phone" field if the given value is not nil.
func (duo *DoctorUpdateOne) SetNillablePhone(s *string) *DoctorUpdateOne {
	if s != nil {
		duo.SetPhone(*s)
	}
	return duo
}

// ClearPhone clears the value of the "phone" field.
func (duo *DoctorUpdateOne) ClearPhone() *DoctorUpdateOne {
	duo.mutation.ClearPhone()
	return duo
}

// SetBio sets the "bio" field.
func (duo *DoctorUpdateOne) SetBio(s string) *DoctorUpdateOne {
	duo.mutation.SetBio(s)
	return duo
}

// SetNillableBio sets the "bio" field if the given value is not nil.
func (duo *DoctorUpdateOne) SetNillableBio(s *string) *DoctorUpdateOne {
	if s != nil {
		duo.SetBio(*s)
	}
	return duo
}

// ClearBio clears the value of the "bio" field.
func (duo *DoctorUpdateOne) ClearBio() *DoctorUpdateOne {
	duo.mutation.ClearBio()
	return duo
}

// SetQualifications sets the "qualifications" field.
func (duo *DoctorUpdateOne) SetQualifications(s []string) *DoctorUpdateOne {
	duo.mutation.SetQualifications(s)
	return duo
}

// AppendQualifications appends s to the "qualifications" field.
func (duo *DoctorUpdateOne) AppendQualifications(s []string) *DoctorUpdateOne {
	duo.mutation.AppendQualifications(s)
	return duo
}

// ClearQualifications clears the value of the "qualifications" field.
func (duo *DoctorUpdateOne) ClearQualifications() *DoctorUpdateOne {
	duo.mutation.ClearQualifications()
	return duo
}

// SetAvailability sets the "availability" field.
func (duo *DoctorUpdateOne) SetAvailability(m map[string]interface{}) *DoctorUpdateOne {
	duo.mutation.SetAvailability(m)
	return duo
}

// SetConsultationDuration sets the "consultation_duration" field.
func (duo *DoctorUpdateOne) SetConsultationDuration(i int) *DoctorUpdateOne {
	duo.mutation.ResetConsultationDuration()
	duo.mutation.SetConsultationDuration(i)
	return duo
}

// SetNillableConsultationDuration sets the "consultation_duration" field if the given value is not nil.
func (duo *DoctorUpdateOne) SetNillableConsultationDuration(i *int) *DoctorUpdateOne {
	if i != nil {
		duo.SetConsultationDuration(*i)
	}
	return duo
}

// AddConsultationDuration adds i to the "consultation_duration" field.
func (duo *DoctorUpdateOne) AddConsultationDuration(i int) *DoctorUpdateOne {
	duo.mutation.AddConsultationDuration(i)
	return duo
}

// SetConsultationFee sets the "consultation_fee" field.
func (duo *DoctorUpdateOne) SetConsultationFee(f float64) *DoctorUpdateOne {
	duo.mutation.ResetConsultationFee()
	duo.mutation.SetConsultationFee(f)
	return duo
}

// SetNillableConsultationFee sets the "consultation_fee" field if the given value is not nil.
func (duo *DoctorUpdateOne) SetNillableConsultationFee(f *float64) *DoctorUpdateOne {
	if f != nil {
		duo.SetConsultationFee(*f)
	}
	return duo
}

// AddConsultationFee adds f to the "consultation_fee" field.
func (duo *DoctorUpdateOne) AddConsultationFee(f float64) *DoctorUpdateOne {
	duo.mutation.AddConsultationFee(f)
	return duo
}

// SetActive sets the "active" field.
func (duo *DoctorUpdateOne) SetActive(b bool) *DoctorUpdateOne {
	duo.mutation.SetActive(b)
	return duo
}

// SetNillableActive sets the "active" field if the given value is not nil.
func (duo *DoctorUpdateOne) SetNillableActive(b *bool) *DoctorUpdateOne {
	if b != nil {
		duo.SetActive(*b)
	}
	return duo
}

// SetCreatedAt sets the "created_at" field.
func (duo *DoctorUpdateOne) SetCreatedAt(t time.Time) *DoctorUpdateOne {
	duo.mutation.SetCreatedAt(t)
	return duo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (duo *DoctorUpdateOne) SetNillableCreatedAt(t *time.Time) *DoctorUpdateOne {
	if t != nil {
		duo.SetCreatedAt(*t)
	}
	return duo
}

// SetUpdatedAt sets the "updated_at" field.
func (duo *DoctorUpdateOne) SetUpdatedAt(t time.Time) *DoctorUpdateOne {
	duo.mutation.SetUpdatedAt(t)
	return duo
}

// SetClinicID sets the "clinic" edge to the Clinic entity by ID.
func (duo *DoctorUpdateOne) SetClinicID(id uuid.UUID) *DoctorUpdateOne {
	duo.mutation.SetClinicID(id)
	return duo
}

// SetNillableClinicID sets the "clinic" edge to the Clinic entity by ID if the given value is not nil.
func (duo *DoctorUpdateOne) SetNillableClinicID(id *uuid.UUID) *DoctorUpdateOne {
	if id != nil {
		duo = duo.SetClinicID(*id)
	}
	return duo
}

// SetClinic sets the "clinic" edge to the Clinic entity.
func (duo *DoctorUpdateOne) SetClinic(c *Clinic) *DoctorUpdateOne {
	return duo.SetClinicID(c.ID)
}

// AddAppointmentIDs adds the "appointments" edge to the Appointment entity by IDs.
func (duo *DoctorUpdateOne) AddAppointmentIDs(ids ...uuid.UUID) *DoctorUpdateOne {
	duo.mutation.AddAppointmentIDs(ids...)
	return duo
}

// AddAppointments adds the "appointments" edges to the Appointment entity.
func (duo *DoctorUpdateOne) AddAppointments(a ...*Appointment) *DoctorUpdateOne {
	ids := make([]uuid.UUID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return duo.AddAppointmentIDs(ids...)
}

// AddScheduleIDs adds the "schedules" edge to the DoctorSchedule entity by IDs.
func (duo *DoctorUpdateOne) AddScheduleIDs(ids ...uuid.UUID) *DoctorUpdateOne {
	duo.mutation.AddScheduleIDs(ids...)
	return duo
}

// AddSchedules adds the "schedules" edges to the DoctorSchedule entity.
func (duo *DoctorUpdateOne) AddSchedules(d ...*DoctorSchedule) *DoctorUpdateOne {
	ids := make([]uuid.UUID, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return duo.AddScheduleIDs(ids...)
}

// Mutation returns the DoctorMutation object of the builder.
func (duo *DoctorUpdateOne) Mutation() *DoctorMutation {
	return duo.mutation
}

// ClearClinic clears the "clinic" edge to the Clinic entity.
func (duo *DoctorUpdateOne) ClearClinic() *DoctorUpdateOne {
	duo.mutation.ClearClinic()
	return duo
}

// ClearAppointments clears all "appointments" edges to the Appointment entity.
func (duo *DoctorUpdateOne) ClearAppointments() *DoctorUpdateOne {
	duo.mutation.ClearAppointments()
	return duo
}

// RemoveAppointmentIDs removes the "appointments" edge to Appointment entities by IDs.
func (duo *DoctorUpdateOne) RemoveAppointmentIDs(ids ...uuid.UUID) *DoctorUpdateOne {
	duo.mutation.RemoveAppointmentIDs(ids...)
	return duo
}

// RemoveAppointments removes "appointments" edges to Appointment entities.
func (duo *DoctorUpdateOne) RemoveAppointments(a ...*Appointment) *DoctorUpdateOne {
	ids := make([]uuid.UUID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return duo.RemoveAppointmentIDs(ids...)
}

// ClearSchedules clears all "schedules" edges to the DoctorSchedule entity.
func (duo *DoctorUpdateOne) ClearSchedules() *DoctorUpdateOne {
	duo.mutation.ClearSchedules()
	return duo
}

// RemoveScheduleIDs removes the "schedules" edge to DoctorSchedule entities by IDs.
func (duo *DoctorUpdateOne) RemoveScheduleIDs(ids ...uuid.UUID) *DoctorUpdateOne {
	duo.mutation.RemoveScheduleIDs(ids...)
	return duo
}

// RemoveSchedules removes "schedules" edges to DoctorSchedule entities.
func (duo *DoctorUpdateOne) RemoveSchedules(d ...*DoctorSchedule) *DoctorUpdateOne {
	ids := make([]uuid.UUID, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return duo.RemoveScheduleIDs(ids...)
}

// Where appends a list predicates to the DoctorUpdate builder.
func (duo *DoctorUpdateOne) Where(ps ...predicate.Doctor) *DoctorUpdateOne {
	duo.mutation.Where(ps...)
	return duo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (duo *DoctorUpdateOne) Select(field string, fields ...string) *DoctorUpdateOne {
	duo.fields = append([]string{field}, fields...)
	return duo
}

// Save executes the query and returns the updated Doctor entity.
func (duo *DoctorUpdateOne) Save(ctx context.Context) (*Doctor, error) {
	duo.defaults()
	return withHooks(ctx, duo.sqlSave, duo.mutation, duo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (duo *DoctorUpdateOne) SaveX(ctx context.Context) *Doctor {
	node, err := duo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (duo *DoctorUpdateOne) Exec(ctx context.Context) error {
	_, err := duo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (duo *DoctorUpdateOne) ExecX(ctx context.Context) {
	if err := duo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (duo *DoctorUpdateOne) defaults() {
	if _, ok := duo.mutation.UpdatedAt(); !ok {
		v := doctor.UpdateDefaultUpdatedAt()
		duo.mutation.SetUpdatedAt(v)
	}
}

func (duo *DoctorUpdateOne) sqlSave(ctx context.Context) (_node *Doctor, err error) {
	_spec := sqlgraph.NewUpdateSpec(doctor.Table, doctor.Columns, sqlgraph.NewFieldSpec(doctor.FieldID, field.TypeUUID))
	id, ok := duo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`db: missing "Doctor.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := duo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, doctor.FieldID)
		for _, f := range fields {
			if !doctor.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("db: invalid field %q for query", f)}
			}
			if f != doctor.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := duo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := duo.mutation.Name(); ok {
		_spec.SetField(doctor.FieldName, field.TypeString, value)
	}
	if value, ok := duo.mutation.Specialization(); ok {
		_spec.SetField(doctor.FieldSpecialization, field.TypeString, value)
	}
	if value, ok := duo.mutation.LicenseNumber(); ok {
		_spec.SetField(doctor.FieldLicenseNumber, field.TypeString, value)
	}
	if duo.mutation.LicenseNumberCleared() {
		_spec.ClearField(doctor.FieldLicenseNumber, field.TypeString)
	}
	if value, ok := duo.mutation.Email(); ok {
		_spec.SetField(doctor.FieldEmail, field.TypeString, value)
	}
	if duo.mutation.EmailCleared() {
		_spec.ClearField(doctor.FieldEmail, field.TypeString)
	}
	if value, ok := duo.mutation.Phone(); ok {
		_spec.SetField(doctor.FieldPhone, field.TypeString, value)
	}
	if duo.mutation.PhoneCleared() {
		_spec.ClearField(doctor.FieldPhone, field.TypeString)
	}
	if value, ok := duo.mutation.Bio(); ok {
		_spec.SetField(doctor.FieldBio, field.TypeString, value)
	}
	if duo.mutation.BioCleared() {
		_spec.ClearField(doctor.FieldBio, field.TypeString)
	}
	if value, ok := duo.mutation.Qualifications(); ok {
		_spec.SetField(doctor.FieldQualifications, field.TypeJSON, value)
	}
	if value, ok := duo.mutation.AppendedQualifications(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, doctor.FieldQualifications, value)
		})
	}
	if duo.mutation.QualificationsCleared() {
		_spec.ClearField(doctor.FieldQualifications, field.TypeJSON)
	}
	if value, ok := duo.mutation.Availability(); ok {
		_spec.SetField(doctor.FieldAvailability, field.TypeJSON, value)
	}
	if value, ok := duo.mutation.ConsultationDuration(); ok {
		_spec.SetField(doctor.FieldConsultationDuration, field.TypeInt, value)
	}
	if value, ok := duo.mutation.AddedConsultationDuration(); ok {
		_spec.AddField(doctor.FieldConsultationDuration, field.TypeInt, value)
	}
	if value, ok := duo.mutation.ConsultationFee(); ok {
		_spec.SetField(doctor.FieldConsultationFee, field.TypeFloat64, value)
	}
	if value, ok := duo.mutation.AddedConsultationFee(); ok {
		_spec.AddField(doctor.FieldConsultationFee, field.TypeFloat64, value)
	}
	if value, ok := duo.mutation.Active(); ok {
		_spec.SetField(doctor.FieldActive, field.TypeBool, value)
	}
	if value, ok := duo.mutation.CreatedAt(); ok {
		_spec.SetField(doctor.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := duo.mutation.UpdatedAt(); ok {
		_spec.SetField(doctor.FieldUpdatedAt, field.TypeTime, value)
	}
	if duo.mutation.ClinicCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   doctor.ClinicTable,
			Columns: []string{doctor.ClinicColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(clinic.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := duo.mutation.ClinicIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   doctor.ClinicTable,
			Columns: []string{doctor.ClinicColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(clinic.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if duo.mutation.AppointmentsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   doctor.AppointmentsTable,
			Columns: []string{doctor.AppointmentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(appointment.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := duo.mutation.RemovedAppointmentsIDs(); len(nodes) > 0 && !duo.mutation.AppointmentsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   doctor.AppointmentsTable,
			Columns: []string{doctor.AppointmentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(appointment.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := duo.mutation.AppointmentsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   doctor.AppointmentsTable,
			Columns: []string{doctor.AppointmentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(appointment.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if duo.mutation.SchedulesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   doctor.SchedulesTable,
			Columns: []string{doctor.SchedulesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(doctorschedule.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := duo.mutation.RemovedSchedulesIDs(); len(nodes) > 0 && !duo.mutation.SchedulesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   doctor.SchedulesTable,
			Columns: []string{doctor.SchedulesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(doctorschedule.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := duo.mutation.SchedulesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   doctor.SchedulesTable,
			Columns: []string{doctor.SchedulesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(doctorschedule.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Doctor{config: duo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, duo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{doctor.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	duo.mutation.done = true
	return _node, nil
}
