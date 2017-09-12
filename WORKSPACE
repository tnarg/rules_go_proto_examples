workspace(name = "com_github_tnarg_rules_go_proto_examples")

git_repository(
    name = "io_bazel_rules_go",
    remote = "https://github.com/bazelbuild/rules_go.git",
    commit = "4be196cc186da9dd396d5a45a3a7f343b6abe2b0",
)

git_repository(
    name = "com_github_tnarg_rules_go_proto",
    remote = "https://github.com/tnarg/rules_go_proto.git",
    commit = "9012b95628c7eb762e7dbd61be93f46ad20b9370",
)

load("@io_bazel_rules_go//go:def.bzl", "go_repositories", "go_repository")
load("@com_github_tnarg_rules_go_proto//go_proto:rules.bzl", "gogo_repositories")

go_repositories()
gogo_repositories()

go_repository(
    name = "com_github_elazarl_go_bindata_assetfs",
    commit = "30f82fa23fd844bd5bb1e5f216db87fd77b5eb43",
    importpath = "github.com/elazarl/go-bindata-assetfs",
)

go_repository(
    name = "com_github_gorilla_context",
    commit = "08b5f424b9271eedf6f9f0ce86cb9396ed337a42",
    importpath = "github.com/gorilla/context",
)

go_repository(
    name = "com_github_gorilla_mux",
    tag = "v1.4.0",
    importpath = "github.com/gorilla/mux",
)

go_repository(
    name = "com_github_google_uuid",
    commit = "7e072fc3a7be179aee6d3359e46015aa8c995314",
    importpath = "github.com/google/uuid",
)
