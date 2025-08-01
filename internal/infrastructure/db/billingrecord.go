// Code generated by ent, DO NOT EDIT.

package db

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/alfariiizi/vandor/internal/infrastructure/db/billingrecord"
	"github.com/alfariiizi/vandor/internal/infrastructure/db/clinic"
	"github.com/alfariiizi/vandor/internal/infrastructure/db/patient"
	"github.com/google/uuid"
)

// BillingRecord is the model entity for the BillingRecord schema.
type BillingRecord struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// InvoiceNumber holds the value of the "invoice_number" field.
	InvoiceNumber string `json:"invoice_number,omitempty"`
	// Amount holds the value of the "amount" field.
	Amount float64 `json:"amount,omitempty"`
	// TaxAmount holds the value of the "tax_amount" field.
	TaxAmount float64 `json:"tax_amount,omitempty"`
	// DiscountAmount holds the value of the "discount_amount" field.
	DiscountAmount float64 `json:"discount_amount,omitempty"`
	// TotalAmount holds the value of the "total_amount" field.
	TotalAmount float64 `json:"total_amount,omitempty"`
	// Currency holds the value of the "currency" field.
	Currency string `json:"currency,omitempty"`
	// PaymentMethod holds the value of the "payment_method" field.
	PaymentMethod billingrecord.PaymentMethod `json:"payment_method,omitempty"`
	// PaymentStatus holds the value of the "payment_status" field.
	PaymentStatus billingrecord.PaymentStatus `json:"payment_status,omitempty"`
	// services and costs breakdown
	LineItems []map[string]interface{} `json:"line_items,omitempty"`
	// DueDate holds the value of the "due_date" field.
	DueDate time.Time `json:"due_date,omitempty"`
	// PaidAt holds the value of the "paid_at" field.
	PaidAt time.Time `json:"paid_at,omitempty"`
	// Notes holds the value of the "notes" field.
	Notes string `json:"notes,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the BillingRecordQuery when eager-loading is set.
	Edges                   BillingRecordEdges `json:"edges"`
	clinic_billing_records  *uuid.UUID
	patient_billing_records *uuid.UUID
	selectValues            sql.SelectValues
}

// BillingRecordEdges holds the relations/edges for other nodes in the graph.
type BillingRecordEdges struct {
	// Clinic holds the value of the clinic edge.
	Clinic *Clinic `json:"clinic,omitempty"`
	// Patient holds the value of the patient edge.
	Patient *Patient `json:"patient,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// ClinicOrErr returns the Clinic value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e BillingRecordEdges) ClinicOrErr() (*Clinic, error) {
	if e.Clinic != nil {
		return e.Clinic, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: clinic.Label}
	}
	return nil, &NotLoadedError{edge: "clinic"}
}

// PatientOrErr returns the Patient value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e BillingRecordEdges) PatientOrErr() (*Patient, error) {
	if e.Patient != nil {
		return e.Patient, nil
	} else if e.loadedTypes[1] {
		return nil, &NotFoundError{label: patient.Label}
	}
	return nil, &NotLoadedError{edge: "patient"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*BillingRecord) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case billingrecord.FieldLineItems:
			values[i] = new([]byte)
		case billingrecord.FieldAmount, billingrecord.FieldTaxAmount, billingrecord.FieldDiscountAmount, billingrecord.FieldTotalAmount:
			values[i] = new(sql.NullFloat64)
		case billingrecord.FieldInvoiceNumber, billingrecord.FieldCurrency, billingrecord.FieldPaymentMethod, billingrecord.FieldPaymentStatus, billingrecord.FieldNotes:
			values[i] = new(sql.NullString)
		case billingrecord.FieldDueDate, billingrecord.FieldPaidAt, billingrecord.FieldCreatedAt, billingrecord.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case billingrecord.FieldID:
			values[i] = new(uuid.UUID)
		case billingrecord.ForeignKeys[0]: // clinic_billing_records
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		case billingrecord.ForeignKeys[1]: // patient_billing_records
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the BillingRecord fields.
func (br *BillingRecord) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case billingrecord.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				br.ID = *value
			}
		case billingrecord.FieldInvoiceNumber:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field invoice_number", values[i])
			} else if value.Valid {
				br.InvoiceNumber = value.String
			}
		case billingrecord.FieldAmount:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field amount", values[i])
			} else if value.Valid {
				br.Amount = value.Float64
			}
		case billingrecord.FieldTaxAmount:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field tax_amount", values[i])
			} else if value.Valid {
				br.TaxAmount = value.Float64
			}
		case billingrecord.FieldDiscountAmount:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field discount_amount", values[i])
			} else if value.Valid {
				br.DiscountAmount = value.Float64
			}
		case billingrecord.FieldTotalAmount:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field total_amount", values[i])
			} else if value.Valid {
				br.TotalAmount = value.Float64
			}
		case billingrecord.FieldCurrency:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field currency", values[i])
			} else if value.Valid {
				br.Currency = value.String
			}
		case billingrecord.FieldPaymentMethod:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field payment_method", values[i])
			} else if value.Valid {
				br.PaymentMethod = billingrecord.PaymentMethod(value.String)
			}
		case billingrecord.FieldPaymentStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field payment_status", values[i])
			} else if value.Valid {
				br.PaymentStatus = billingrecord.PaymentStatus(value.String)
			}
		case billingrecord.FieldLineItems:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field line_items", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &br.LineItems); err != nil {
					return fmt.Errorf("unmarshal field line_items: %w", err)
				}
			}
		case billingrecord.FieldDueDate:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field due_date", values[i])
			} else if value.Valid {
				br.DueDate = value.Time
			}
		case billingrecord.FieldPaidAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field paid_at", values[i])
			} else if value.Valid {
				br.PaidAt = value.Time
			}
		case billingrecord.FieldNotes:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field notes", values[i])
			} else if value.Valid {
				br.Notes = value.String
			}
		case billingrecord.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				br.CreatedAt = value.Time
			}
		case billingrecord.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				br.UpdatedAt = value.Time
			}
		case billingrecord.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field clinic_billing_records", values[i])
			} else if value.Valid {
				br.clinic_billing_records = new(uuid.UUID)
				*br.clinic_billing_records = *value.S.(*uuid.UUID)
			}
		case billingrecord.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field patient_billing_records", values[i])
			} else if value.Valid {
				br.patient_billing_records = new(uuid.UUID)
				*br.patient_billing_records = *value.S.(*uuid.UUID)
			}
		default:
			br.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the BillingRecord.
// This includes values selected through modifiers, order, etc.
func (br *BillingRecord) Value(name string) (ent.Value, error) {
	return br.selectValues.Get(name)
}

// QueryClinic queries the "clinic" edge of the BillingRecord entity.
func (br *BillingRecord) QueryClinic() *ClinicQuery {
	return NewBillingRecordClient(br.config).QueryClinic(br)
}

// QueryPatient queries the "patient" edge of the BillingRecord entity.
func (br *BillingRecord) QueryPatient() *PatientQuery {
	return NewBillingRecordClient(br.config).QueryPatient(br)
}

// Update returns a builder for updating this BillingRecord.
// Note that you need to call BillingRecord.Unwrap() before calling this method if this BillingRecord
// was returned from a transaction, and the transaction was committed or rolled back.
func (br *BillingRecord) Update() *BillingRecordUpdateOne {
	return NewBillingRecordClient(br.config).UpdateOne(br)
}

// Unwrap unwraps the BillingRecord entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (br *BillingRecord) Unwrap() *BillingRecord {
	_tx, ok := br.config.driver.(*txDriver)
	if !ok {
		panic("db: BillingRecord is not a transactional entity")
	}
	br.config.driver = _tx.drv
	return br
}

// String implements the fmt.Stringer.
func (br *BillingRecord) String() string {
	var builder strings.Builder
	builder.WriteString("BillingRecord(")
	builder.WriteString(fmt.Sprintf("id=%v, ", br.ID))
	builder.WriteString("invoice_number=")
	builder.WriteString(br.InvoiceNumber)
	builder.WriteString(", ")
	builder.WriteString("amount=")
	builder.WriteString(fmt.Sprintf("%v", br.Amount))
	builder.WriteString(", ")
	builder.WriteString("tax_amount=")
	builder.WriteString(fmt.Sprintf("%v", br.TaxAmount))
	builder.WriteString(", ")
	builder.WriteString("discount_amount=")
	builder.WriteString(fmt.Sprintf("%v", br.DiscountAmount))
	builder.WriteString(", ")
	builder.WriteString("total_amount=")
	builder.WriteString(fmt.Sprintf("%v", br.TotalAmount))
	builder.WriteString(", ")
	builder.WriteString("currency=")
	builder.WriteString(br.Currency)
	builder.WriteString(", ")
	builder.WriteString("payment_method=")
	builder.WriteString(fmt.Sprintf("%v", br.PaymentMethod))
	builder.WriteString(", ")
	builder.WriteString("payment_status=")
	builder.WriteString(fmt.Sprintf("%v", br.PaymentStatus))
	builder.WriteString(", ")
	builder.WriteString("line_items=")
	builder.WriteString(fmt.Sprintf("%v", br.LineItems))
	builder.WriteString(", ")
	builder.WriteString("due_date=")
	builder.WriteString(br.DueDate.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("paid_at=")
	builder.WriteString(br.PaidAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("notes=")
	builder.WriteString(br.Notes)
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(br.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(br.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// BillingRecords is a parsable slice of BillingRecord.
type BillingRecords []*BillingRecord
