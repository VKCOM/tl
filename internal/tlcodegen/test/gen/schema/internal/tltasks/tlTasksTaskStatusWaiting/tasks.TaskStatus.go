// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
package tlTasksTaskStatusWaiting

import (
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/schema/internal"
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/schema/internal/tltasks/tlTasksTaskStatus"
	"github.com/vkcom/tl/pkg/basictl"
)

var _ = basictl.NatWrite
var _ = internal.ErrorInvalidEnumTag

func TasksTaskStatusWaiting() tlTasksTaskStatus.TasksTaskStatus {
	return tlTasksTaskStatus.TasksTaskStatus__MakeEnum(2)
}
