package utils

const (
	CLOSEDSTATUS = iota + 1 //已关闭
	NOTPROCESSED            //未处理
	PROCESSING              //处理中
	REJECTED                //已拒绝
	COMPLETED               //已完成
	PENDING                 //已挂起
	REOPEN                  //重新打开
)

var CreateMap = map[int32]bool{
	PROCESSING: true,
	REJECTED:   true,
	COMPLETED:  true,
	PENDING:    true,
}

var ProcessorMap = map[int32]bool{
	CLOSEDSTATUS: true,
	NOTPROCESSED: true,
	REOPEN:       true,
}

const NOSPECIFIED = "未指定"

const POSFILENAME = "pos.txt"

const UPDATE = "update"
const INSERT = "insert"
const DELETE = "delete"

const (
	DeploymentResource = "deployment"
	JobResource        = "job"
	JobCornResource    = "jobcorn"
)
