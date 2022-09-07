package utils

import (
    "github.com/bwmarrin/snowflake"
)

var uuidNode *snowflake.Node

func InitUuid(nodeInstance *snowflake.Node)  {
    uuidNode = nodeInstance
}

func GetUuid() *snowflake.Node  {
    return uuidNode
}
