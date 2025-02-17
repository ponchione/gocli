package tests

//func TestRoll_Success(t *testing.T) {
//	args := []string{"1", "10"}
//	err := cmd.Roll(args)
//	if err != nil {
//		t.Errorf("Roll() resulted in error: %v", err)
//	}
//}
//
//func TestRoll_InvalidArguments(t *testing.T) {
//	args := []string{"a", "b"}
//	err := cmd.Roll(args)
//	if err == nil {
//		t.Errorf("Expected an error due to invalid arguments, but got none")
//	}
//}
//
//func TestRoll_InsufficientArguments(t *testing.T) {
//	args := []string{"5"}
//	err := cmd.Roll(args)
//	if err == nil {
//		t.Errorf("Expected an error due to insufficient arguments, but got none")
//	}
//}
//
//func TestRoll_ExcessArguments(t *testing.T) {
//	args := []string{"1", "10", "20"}
//	err := cmd.Execute(args)
//	if err == nil {
//		t.Errorf("Expected an error due to excess arguments, but got none")
//	}
//}
//
//func TestRoll_NegativeNumbers(t *testing.T) {
//	args := []string{"-10", "-1"}
//	err := cmd.Roll(args)
//	if err != nil {
//		t.Errorf("Roll() resulted in error for negative numbers: %v", err)
//	}
//}
//
//func TestRoll_InvalidRange(t *testing.T) {
//	args := []string{"10", "1"}
//	err := cmd.Roll(args)
//	if err == nil {
//		t.Errorf("Expected an error for min > max, but got none")
//	}
//}
//
//func TestValidateNumberArgs_Empty(t *testing.T) {
//	args := []string{}
//	err := cmd.validateNumberArgs(args)
//	if err == nil {
//		t.Errorf("Expected error for empty arguments, but got none")
//	}
//}
//
//func TestValidateNumberArgs_TooMany(t *testing.T) {
//	args := []string{"1", "2", "3"}
//	err := cmd.validateNumberArgs(args)
//	if err == nil {
//		t.Errorf("Expected error for too many arguments, but got none")
//	}
//}
