package test

import (
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"testing"

	_ "github.com/lib/pq"
)

var ATM ArrtestModel

func TestNewArrtestModel(t *testing.T) {

	ATM = NewArrtestModel(sqlx.NewSqlConn("postgres", "postgres://postgres:123456@127.0.0.1:5432/sc?sslmode=disable"))
	//ATM = NewArrtestModel(sqlx.NewSqlConn("postgres", "user=postgres password=123456 dbname=sc sslmode=disable host=127.0.0.1 port=5432"))

	find, err := ATM.TestFind(1)
	if err != nil {
		return
	}

	fmt.Printf("%+v", find)

	//arr1 := []int64{1, 2, 3}
	//arr2 := []string{"123", "嘻嘻嘻"}
	//arr3 := []int64{1111, 2222}
	//_, err := ATM.TestInsert(&Arrtest{
	//	Arr1: strings.Join(arr1, ""),
	//	Arr2: arr2,
	//	Arr3: arr3,
	//})
	//if err != nil {
	//	return
	//}

}
