package repository

import (
	"database/sql"
	"jakpat-test-2/entity"
	"reflect"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
)

func TestDBPostgres_CreateUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal("Error connection to DB")
	}
	defer db.Close()
	sqlxDB := sqlx.NewDb(db, "sqlmock")

	type fields struct {
		DB *sqlx.DB
	}
	type args struct {
		input entity.Users
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func()
		want    int
		wantErr bool
	}{
		{
			name: "error",
			fields: fields{
				DB: sqlxDB,
			},
			args: args{
				input: entity.Users{
					Name:     "asd",
					Password: "qwe",
					Role:     1,
					Status:   1,
				},
			},
			mock: func() {
				mock.ExpectQuery(regexp.QuoteMeta("INSERT INTO users (name, password, role, status) VALUES ($1, $2, $3, $4) RETURNING id")).
					WithArgs("asd", "qwe", int(1), int(1)).WillReturnError(sql.ErrNoRows)
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "success",
			fields: fields{
				DB: sqlxDB,
			},
			args: args{
				input: entity.Users{
					Name:     "asd",
					Password: "qwe",
					Role:     1,
					Status:   1,
				},
			},
			mock: func() {
				mock.ExpectQuery(regexp.QuoteMeta("INSERT INTO users (name, password, role, status) VALUES ($1, $2, $3, $4) RETURNING id")).
					WithArgs("asd", "qwe", int(1), int(1)).WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
			},
			want:    1,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			r := NewDBPostgres(tt.fields.DB)
			got, err := r.CreateUser(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("DBPostgres.CreateUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("DBPostgres.CreateUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDBPostgres_GetUserByIdAndStatus(t *testing.T) {
	mockTime := time.Now()

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal("Error connection to DB")
	}
	defer db.Close()
	sqlxDB := sqlx.NewDb(db, "sqlmock")

	type fields struct {
		DB *sqlx.DB
	}
	type args struct {
		id     int
		status int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func()
		want    entity.Users
		wantErr bool
	}{
		{
			name: "error",
			fields: fields{
				DB: sqlxDB,
			},
			args: args{
				id:     1,
				status: 1,
			},
			mock: func() {
				mock.ExpectQuery(regexp.QuoteMeta("SELECT id, name, role, join_date FROM users WHERE id = $1 AND status = $2")).
					WithArgs(int(1), int(1)).WillReturnError(sql.ErrNoRows)
			},
			want:    entity.Users{},
			wantErr: true,
		},
		{
			name: "success",
			fields: fields{
				DB: sqlxDB,
			},
			args: args{
				id:     1,
				status: 1,
			},
			mock: func() {
				mock.ExpectQuery(regexp.QuoteMeta("SELECT id, name, role, join_date FROM users WHERE id = $1 AND status = $2")).
					WithArgs(int(1), int(1)).WillReturnRows(mock.NewRows([]string{"id", "name", "role", "join_date"}).AddRow(1, "asd", 1, mockTime))
			},
			want: entity.Users{
				ID:       1,
				Name:     "asd",
				Role:     1,
				JoinDate: mockTime,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			r := NewDBPostgres(tt.fields.DB)
			got, err := r.GetUserByIdAndStatus(tt.args.id, tt.args.status)
			if (err != nil) != tt.wantErr {
				t.Errorf("DBPostgres.GetUserByIdAndStatus() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DBPostgres.GetUserByIdAndStatus() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDBPostgres_GetUserByNameAndPassword(t *testing.T) {
	mockTime := time.Now()

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal("Error connection to DB")
	}
	defer db.Close()
	sqlxDB := sqlx.NewDb(db, "sqlmock")

	type fields struct {
		DB *sqlx.DB
	}
	type args struct {
		name     string
		password string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func()
		want    entity.Users
		wantErr bool
	}{
		{
			name: "error",
			fields: fields{
				DB: sqlxDB,
			},
			args: args{
				name:     "asd",
				password: "qwe",
			},
			mock: func() {
				mock.ExpectQuery(regexp.QuoteMeta("SELECT id, name, role, join_date, status FROM users WHERE name = $1 AND password = $2 AND status = 1")).
					WithArgs("asd", "qwe").WillReturnError(sql.ErrNoRows)
			},
			want:    entity.Users{},
			wantErr: true,
		},
		{
			name: "success",
			fields: fields{
				DB: sqlxDB,
			},
			args: args{
				name:     "asd",
				password: "qwe",
			},
			mock: func() {
				mock.ExpectQuery(regexp.QuoteMeta("SELECT id, name, role, join_date, status FROM users WHERE name = $1 AND password = $2 AND status = 1")).
					WithArgs("asd", "qwe").WillReturnRows(mock.NewRows([]string{"id", "name", "role", "join_date"}).AddRow(1, "asd", 1, mockTime))
			},
			want: entity.Users{
				ID:       1,
				Name:     "asd",
				Role:     1,
				JoinDate: mockTime,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			r := NewDBPostgres(tt.fields.DB)
			got, err := r.GetUserByNameAndPassword(tt.args.name, tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("DBPostgres.GetUserByNameAndPassword() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DBPostgres.GetUserByNameAndPassword() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDBPostgres_AddItem(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal("Error connection to DB")
	}
	defer db.Close()
	sqlxDB := sqlx.NewDb(db, "sqlmock")

	type fields struct {
		DB *sqlx.DB
	}
	type args struct {
		input entity.Items
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func()
		want    int
		wantErr bool
	}{
		{
			name: "error",
			fields: fields{
				DB: sqlxDB,
			},
			args: args{
				input: entity.Items{
					SellerID: 3,
					Name:     "asd",
					Stock:    2,
					Status:   1,
				},
			},
			mock: func() {
				mock.ExpectQuery(regexp.QuoteMeta("INSERT INTO items (seller_id, name, stock, status) VALUES ($1, $2, $3, $4) RETURNING id")).
					WithArgs(int(3), "asd", int(2), int(1)).WillReturnError(sql.ErrNoRows)
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "success",
			fields: fields{
				DB: sqlxDB,
			},
			args: args{
				input: entity.Items{
					SellerID: 3,
					Name:     "asd",
					Stock:    2,
					Status:   1,
				},
			},
			mock: func() {
				mock.ExpectQuery(regexp.QuoteMeta("INSERT INTO items (seller_id, name, stock, status) VALUES ($1, $2, $3, $4) RETURNING id")).
					WithArgs(int(3), "asd", int(2), int(1)).WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
			},
			want:    1,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			r := NewDBPostgres(tt.fields.DB)
			got, err := r.AddItem(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("DBPostgres.AddItem() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("DBPostgres.AddItem() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDBPostgres_GetItemByIdAndStatus(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal("Error connection to DB")
	}
	defer db.Close()
	sqlxDB := sqlx.NewDb(db, "sqlmock")

	type fields struct {
		DB *sqlx.DB
	}
	type args struct {
		id     int
		status int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func()
		want    entity.Items
		wantErr bool
	}{
		{
			name: "error",
			fields: fields{
				DB: sqlxDB,
			},
			args: args{
				id:     1,
				status: 1,
			},
			mock: func() {
				mock.ExpectQuery(regexp.QuoteMeta("SELECT id, seller_id, name, stock FROM items WHERE id = $1 AND status = $2")).
					WithArgs(int(1), int(1)).WillReturnError(sql.ErrNoRows)
			},
			want:    entity.Items{},
			wantErr: true,
		},
		{
			name: "success",
			fields: fields{
				DB: sqlxDB,
			},
			args: args{
				id:     1,
				status: 1,
			},
			mock: func() {
				mock.ExpectQuery(regexp.QuoteMeta("SELECT id, seller_id, name, stock FROM items WHERE id = $1 AND status = $2")).
					WithArgs(int(1), int(1)).WillReturnRows(mock.NewRows([]string{"id", "seller_id", "name", "stock"}).AddRow(1, 2, "asd", 4))
			},
			want: entity.Items{
				ID:       1,
				SellerID: 2,
				Name:     "asd",
				Stock:    4,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			r := NewDBPostgres(tt.fields.DB)
			got, err := r.GetItemByIdAndStatus(tt.args.id, tt.args.status)
			if (err != nil) != tt.wantErr {
				t.Errorf("DBPostgres.GetItemByIdAndStatus() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DBPostgres.GetItemByIdAndStatus() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDBPostgres_UpdateItemById(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal("Error connection to DB")
	}
	defer db.Close()
	sqlxDB := sqlx.NewDb(db, "sqlmock")

	type fields struct {
		DB *sqlx.DB
	}
	type args struct {
		id    int
		input entity.Items
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func()
		wantErr bool
	}{
		{
			name: "error",
			fields: fields{
				DB: sqlxDB,
			},
			args: args{
				id: 1,
				input: entity.Items{
					SellerID: 2,
					Name:     "asd",
					Stock:    4,
					Status:   1,
				},
			},
			mock: func() {
				mock.ExpectExec(regexp.QuoteMeta("UPDATE items SET seller_id = $1, name = $2, stock = $3, status = $4 WHERE id = $5")).
					WithArgs(int(2), "asd", int(4), int(1), int(1)).WillReturnError(sql.ErrNoRows)
			},
			wantErr: true,
		},
		{
			name: "success",
			fields: fields{
				DB: sqlxDB,
			},
			args: args{
				id: 1,
				input: entity.Items{
					SellerID: 2,
					Name:     "asd",
					Stock:    4,
					Status:   1,
				},
			},
			mock: func() {
				mock.ExpectExec(regexp.QuoteMeta("UPDATE items SET seller_id = $1, name = $2, stock = $3, status = $4 WHERE id = $5")).
					WithArgs(int(2), "asd", int(4), int(1), int(1)).WillReturnResult(sqlmock.NewResult(1, 1))
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			r := NewDBPostgres(tt.fields.DB)
			if err := r.UpdateItemById(tt.args.id, tt.args.input); (err != nil) != tt.wantErr {
				t.Errorf("DBPostgres.UpdateItemById() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDBPostgres_GetItemsBySellerIdAndStatus(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal("Error connection to DB")
	}
	defer db.Close()
	sqlxDB := sqlx.NewDb(db, "sqlmock")

	type fields struct {
		DB *sqlx.DB
	}
	type args struct {
		sellerID int
		status   int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func()
		want    []entity.Items
		wantErr bool
	}{
		{
			name: "error",
			fields: fields{
				DB: sqlxDB,
			},
			args: args{
				sellerID: 1,
				status:   1,
			},
			mock: func() {
				mock.ExpectQuery(regexp.QuoteMeta("SELECT id, seller_id, name, stock, status FROM items WHERE seller_id = $1 AND status = $2")).
					WithArgs(int(1), int(1)).WillReturnError(sql.ErrNoRows)
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "success",
			fields: fields{
				DB: sqlxDB,
			},
			args: args{
				sellerID: 1,
				status:   1,
			},
			mock: func() {
				mock.ExpectQuery(regexp.QuoteMeta("SELECT id, seller_id, name, stock, status FROM items WHERE seller_id = $1 AND status = $2")).
					WithArgs(int(1), int(1)).WillReturnRows(mock.NewRows([]string{"id", "seller_id", "name", "stock"}).AddRow(1, 2, "asd", 4))
			},
			want: []entity.Items{
				{
					ID:       1,
					SellerID: 2,
					Name:     "asd",
					Stock:    4,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			r := NewDBPostgres(tt.fields.DB)
			got, err := r.GetItemsBySellerIdAndStatus(tt.args.sellerID, tt.args.status)
			if (err != nil) != tt.wantErr {
				t.Errorf("DBPostgres.GetItemsBySellerIdAndStatus() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DBPostgres.GetItemsBySellerIdAndStatus() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDBPostgres_CreateOrder(t *testing.T) {
	mockTime := time.Now()

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal("Error connection to DB")
	}
	defer db.Close()
	sqlxDB := sqlx.NewDb(db, "sqlmock")

	type fields struct {
		DB *sqlx.DB
	}
	type args struct {
		input entity.Order
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func()
		want    int
		wantErr bool
	}{
		{
			name: "error",
			fields: fields{
				DB: sqlxDB,
			},
			args: args{
				input: entity.Order{
					ItemID:      1,
					BuyerID:     2,
					Status:      1,
					ExpiredDate: mockTime,
				},
			},
			mock: func() {
				mock.ExpectQuery(regexp.QuoteMeta("INSERT INTO orders (item_id, buyer_id, status, expired_date) VALUES ($1, $2, $3, $4) RETURNING id")).
					WithArgs(int(1), int(2), int(1), mockTime).WillReturnError(sql.ErrNoRows)
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "success",
			fields: fields{
				DB: sqlxDB,
			},
			args: args{
				input: entity.Order{
					ItemID:      1,
					BuyerID:     2,
					Status:      1,
					ExpiredDate: mockTime,
				},
			},
			mock: func() {
				mock.ExpectQuery(regexp.QuoteMeta("INSERT INTO orders (item_id, buyer_id, status, expired_date) VALUES ($1, $2, $3, $4) RETURNING id")).
					WithArgs(int(1), int(2), int(1), mockTime).WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
			},
			want:    1,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			r := NewDBPostgres(tt.fields.DB)
			got, err := r.CreateOrder(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("DBPostgres.CreateOrder() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("DBPostgres.CreateOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDBPostgres_GetOrderById(t *testing.T) {
	mockTime := time.Now()

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal("Error connection to DB")
	}
	defer db.Close()
	sqlxDB := sqlx.NewDb(db, "sqlmock")

	type fields struct {
		DB *sqlx.DB
	}
	type args struct {
		id int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func()
		want    entity.Order
		wantErr bool
	}{
		{
			name: "error",
			fields: fields{
				DB: sqlxDB,
			},
			args: args{
				id: 1,
			},
			mock: func() {
				mock.ExpectQuery(regexp.QuoteMeta(`
				SELECT orders.id, orders.item_id, orders.buyer_id, items.seller_id, orders.status, orders.expired_date, orders.created_date, orders.last_updated 
				FROM orders
				JOIN items on orders.item_id = items.id
				WHERE orders.id = $1`)).
					WithArgs(int(1)).WillReturnError(sql.ErrNoRows)
			},
			want:    entity.Order{},
			wantErr: true,
		},
		{
			name: "success",
			fields: fields{
				DB: sqlxDB,
			},
			args: args{
				id: 1,
			},
			mock: func() {
				mock.ExpectQuery(regexp.QuoteMeta(`
				SELECT orders.id, orders.item_id, orders.buyer_id, items.seller_id, orders.status, orders.expired_date, orders.created_date, orders.last_updated 
				FROM orders
				JOIN items on orders.item_id = items.id
				WHERE orders.id = $1`)).
					WithArgs(int(1)).WillReturnRows(sqlmock.NewRows([]string{"id", "item_id", "buyer_id", "seller_id", "status", "expired_date", "created_date", "last_updated"}).AddRow(1, 2, 3, 4, 5, mockTime, mockTime, mockTime))
			},
			want: entity.Order{
				ID:          1,
				ItemID:      2,
				BuyerID:     3,
				SellerID:    4,
				Status:      5,
				ExpiredDate: mockTime,
				CreatedDate: mockTime,
				LastUpdated: mockTime,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			r := NewDBPostgres(tt.fields.DB)
			got, err := r.GetOrderById(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("DBPostgres.GetOrderById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DBPostgres.GetOrderById() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDBPostgres_GetOrdersBySellerId(t *testing.T) {
	mockTime := time.Now()

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal("Error connection to DB")
	}
	defer db.Close()
	sqlxDB := sqlx.NewDb(db, "sqlmock")

	type fields struct {
		DB *sqlx.DB
	}
	type args struct {
		sellerID int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func()
		want    []entity.Order
		wantErr bool
	}{
		{
			name: "error",
			fields: fields{
				DB: sqlxDB,
			},
			args: args{
				sellerID: 1,
			},
			mock: func() {
				mock.ExpectQuery(regexp.QuoteMeta(`
				SELECT orders.id, orders.item_id, orders.buyer_id, items.seller_id, orders.status, orders.expired_date, orders.created_date, orders.last_updated 
				FROM orders
				JOIN items on orders.item_id = items.id
				WHERE items.seller_id = $1`)).
					WithArgs(int(1)).WillReturnError(sql.ErrNoRows)
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "success",
			fields: fields{
				DB: sqlxDB,
			},
			args: args{
				sellerID: 1,
			},
			mock: func() {
				mock.ExpectQuery(regexp.QuoteMeta(`
				SELECT orders.id, orders.item_id, orders.buyer_id, items.seller_id, orders.status, orders.expired_date, orders.created_date, orders.last_updated 
				FROM orders
				JOIN items on orders.item_id = items.id
				WHERE items.seller_id = $1`)).
					WithArgs(int(1)).WillReturnRows(sqlmock.NewRows([]string{"id", "item_id", "buyer_id", "seller_id", "status", "expired_date", "created_date", "last_updated"}).AddRow(1, 2, 3, 4, 5, mockTime, mockTime, mockTime))
			},
			want: []entity.Order{
				{
					ID:          1,
					ItemID:      2,
					BuyerID:     3,
					SellerID:    4,
					Status:      5,
					ExpiredDate: mockTime,
					CreatedDate: mockTime,
					LastUpdated: mockTime,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			r := NewDBPostgres(tt.fields.DB)
			got, err := r.GetOrdersBySellerId(tt.args.sellerID)
			if (err != nil) != tt.wantErr {
				t.Errorf("DBPostgres.GetOrdersBySellerId() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DBPostgres.GetOrdersBySellerId() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDBPostgres_UpdateOrderById(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal("Error connection to DB")
	}
	defer db.Close()
	sqlxDB := sqlx.NewDb(db, "sqlmock")

	type fields struct {
		DB *sqlx.DB
	}
	type args struct {
		id    int
		input entity.Order
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func()
		wantErr bool
	}{
		{
			name: "error",
			fields: fields{
				DB: sqlxDB,
			},
			args: args{
				id: 1,
				input: entity.Order{
					Status: 2,
				},
			},
			mock: func() {
				mock.ExpectExec(regexp.QuoteMeta("UPDATE orders SET status = $1, last_updated = NOW() WHERE id = $2")).
					WithArgs(int(2), int(1)).WillReturnError(sql.ErrNoRows)
			},
			wantErr: true,
		},
		{
			name: "success",
			fields: fields{
				DB: sqlxDB,
			},
			args: args{
				id: 1,
				input: entity.Order{
					Status: 2,
				},
			},
			mock: func() {
				mock.ExpectExec(regexp.QuoteMeta("UPDATE orders SET status = $1, last_updated = NOW() WHERE id = $2")).
					WithArgs(int(2), int(1)).WillReturnResult(sqlmock.NewResult(1, 1))
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			r := NewDBPostgres(tt.fields.DB)
			if err := r.UpdateOrderById(tt.args.id, tt.args.input); (err != nil) != tt.wantErr {
				t.Errorf("DBPostgres.UpdateOrderById() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
