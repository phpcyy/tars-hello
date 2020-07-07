APP       := App
TARGET    := Hello
MFLAGS    :=
DFLAGS    :=
CONFIG    :=
STRIP_FLAG:= N
J2GO_FLAG:= 

libpath=${subst :, ,$(GOPATH)}
$(foreach path,$(libpath),$(eval -include vendor/github.com/TarsCloud/TarsGo/tars/makefile.tars))