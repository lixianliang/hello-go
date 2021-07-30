package main

import (
	"fmt"
)

type Machine struct {
	state IState
}

func (m *Machine) SetState(state IState) {
	m.state = state
}

func (m *Machine) GetStateName() string {
	return m.state.GetName()
}

func (m *Machine) Approval() {
	m.state.Approval(m)
}

func (m *Machine) Reject() {
	m.state.Reject(m)
}

type IState interface {
	Approval(m *Machine)
	Reject(m *Machine)
	GetName() string
}

type leaderApproveState struct{}

func (leaderApproveState) Approval(m *Machine) {
	fmt.Println("leader 审批成功")
	m.SetState(GetFinanceApproveState())
}

func (leaderApproveState) GetName() string {
	return "leaderApproveState"
}

func (leaderApproveState) Reject(m *Machine) {}

func GetLeaderApproveState() IState {
	return &leaderApproveState{}
}

type financeApproveState struct{}

func (f financeApproveState) Approval(m *Machine) {
	fmt.Println("财务审批成功")
	fmt.Println("出发打款操作")
}

func (f financeApproveState) Reject(m *Machine) {
	m.SetState(GetLeaderApproveState())
}

func (f financeApproveState) GetName() string {
	return "financeApproveState"
}

func GetFinanceApproveState() IState {
	return &financeApproveState{}
}

func main() {
	m := &Machine{state: GetLeaderApproveState()}
	fmt.Println(m.GetStateName())
	m.Approval()
	fmt.Println(m.GetStateName())

	m.Reject()
	fmt.Println(m.GetStateName())
	m.Approval()
	fmt.Println(m.GetStateName())

	m.Approval()

}
