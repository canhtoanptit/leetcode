package sodility

import "testing"

func Test_solutionA(t *testing.T) {
	type args struct {
		id            int
		employeeInfos []*EmployeeInfo
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "tc1",
			args: args{
				id: 1,
				employeeInfos: []*EmployeeInfo{
					{
						Id:             1,
						ImportantValue: 5,
						SubIDS:         []int{2, 3},
					},
					{
						Id:             2,
						ImportantValue: 3,
						SubIDS:         []int{},
					},
					{
						Id:             3,
						ImportantValue: 3,
						SubIDS:         []int{},
					},
				},
			},
			want: 11,
		},
		{
			name: "tc2",
			args: args{
				id: 1,
				employeeInfos: []*EmployeeInfo{
					{
						Id:             1,
						ImportantValue: 5,
						SubIDS:         []int{2},
					},
					{
						Id:             2,
						ImportantValue: 3,
						SubIDS:         []int{3},
					},
					{
						Id:             3,
						ImportantValue: 3,
						SubIDS:         []int{},
					},
				},
			},
			want: 11,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solutionA(tt.args.id, tt.args.employeeInfos); got != tt.want {
				t.Errorf("solutionA() = %v, want %v", got, tt.want)
			}
		})
	}
}
