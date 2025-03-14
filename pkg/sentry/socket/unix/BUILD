load("//tools:defs.bzl", "go_library")
load("//tools/go_generics:defs.bzl", "go_template_instance")

package(
    default_applicable_licenses = ["//:license"],
    licenses = ["notice"],
)

go_template_instance(
    name = "socket_refs",
    out = "socket_refs.go",
    package = "unix",
    prefix = "socket",
    template = "//pkg/refs:refs_template",
    types = {
        "T": "Socket",
    },
)

go_library(
    name = "unix",
    srcs = [
        "io.go",
        "socket_refs.go",
        "unix.go",
    ],
    visibility = ["//pkg/sentry:internal"],
    deps = [
        "//pkg/abi/linux",
        "//pkg/atomicbitops",
        "//pkg/context",
        "//pkg/errors/linuxerr",
        "//pkg/fspath",
        "//pkg/hostarch",
        "//pkg/log",
        "//pkg/marshal",
        "//pkg/refs",
        "//pkg/safemem",
        "//pkg/sentry/arch",
        "//pkg/sentry/fsimpl/lock",
        "//pkg/sentry/fsimpl/sockfs",
        "//pkg/sentry/fsutil",
        "//pkg/sentry/inet",
        "//pkg/sentry/kernel",
        "//pkg/sentry/ktime",
        "//pkg/sentry/socket",
        "//pkg/sentry/socket/control",
        "//pkg/sentry/socket/netstack",
        "//pkg/sentry/socket/unix/transport",
        "//pkg/sentry/vfs",
        "//pkg/syserr",
        "//pkg/tcpip",
        "//pkg/usermem",
        "//pkg/waiter",
        "@org_golang_x_sys//unix:go_default_library",
    ],
)
