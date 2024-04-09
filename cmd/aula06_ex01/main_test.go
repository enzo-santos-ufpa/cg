package main

import (
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
)

func TestRun(t *testing.T) {
	var builder strings.Builder
	err := run(&builder)
	require.Nil(t, err)
	require.Equal(t, `⁴ ┊  ┃  ┊┊ 
³╌┊╌╌┃╌╌░█╌
² ┊ ░░░░┊░ 
¹╌█░╌┃╌╌┊░╌
⁰━━░░╋━━━░━
¹ ┊  ░  ░┊ 
² ┊  ┃░░░┊ 
³╌┊╌╌┃╌╌█┊╌
⁴ ┊  ┃  ┊┊ 
 ⁴³²¹⁰¹²³⁴⁵

⁷ ┃     ┊    ┊ 
⁶╌┃╌╌╌╌╌█░╌╌╌┊╌
⁵ ┃   ░░┊ ░░ ┊ 
⁴╌┃╌░░╌╌░░░░░█╌
³╌█░░░░░┊╌╌╌╌┊╌
² ┃     ┊    ┊ 
 ¹⁰¹²³⁴⁵⁶⁷⁸⁹⁰¹²`, builder.String())
}
