# This is for the Repository Test Transaction RollBack

Do not want to write a code every single time when I create a new Project.

## How To Use

```go

func Test_Rollback(t *testing.T) {
	testDB := NewTestDB(&gorm.DB{})
    
	testDB.RunWithRollback(func(tx *gorm.DB) error) error {
        // Do something with tx
        return nil
    })
}
```