load("@bazel_gazelle//:deps.bzl", "go_repository")

def go_dependencies():
    go_repository(
        name = "co_honnef_go_tools",
        build_file_proto_mode = "disable",
        importpath = "honnef.co/go/tools",
        sum = "h1:UoveltGrhghAA7ePc+e+QYDHXrBps2PqFZiHkGR/xK8=",
        version = "v0.0.1-2020.1.4",
    )
    go_repository(
        name = "com_github_99designs_gqlgen",
        build_file_proto_mode = "disable",
        importpath = "github.com/99designs/gqlgen",
        sum = "h1:nGGP+sUJ6D3guzjVBgoH1PrZxoU4lUdfR/Q8THYrAJI=",
        version = "v0.17.5-0.20220428154617-9250f9ac1f90",
    )

    go_repository(
        name = "com_github_afex_hystrix_go",
        build_file_proto_mode = "disable",
        importpath = "github.com/afex/hystrix-go",
        sum = "h1:rFw4nCn9iMW+Vajsk51NtYIcwSTkXr+JGrMd36kTDJw=",
        version = "v0.0.0-20180502004556-fa1af6a1f4f5",
    )
    go_repository(
        name = "com_github_agext_levenshtein",
        build_file_proto_mode = "disable",
        importpath = "github.com/agext/levenshtein",
        sum = "h1:YB2fHEn0UJagG8T1rrWknE3ZQzWM06O8AMAatNn7lmo=",
        version = "v1.2.3",
    )
    go_repository(
        name = "com_github_agiledragon_gomonkey",
        build_file_proto_mode = "disable",
        importpath = "github.com/agiledragon/gomonkey",
        sum = "h1:eXKi9/piiC3cjJD1658mEE2o3NjkJ5vDLgYjCQu0Xlw=",
        version = "v2.0.2+incompatible",
    )
    go_repository(
        name = "com_github_agnivade_levenshtein",
        build_file_proto_mode = "disable",
        importpath = "github.com/agnivade/levenshtein",
        sum = "h1:QY8M92nrzkmr798gCo3kmMyqXFzdQVpxLlGPRBij0P8=",
        version = "v1.1.1",
    )

    go_repository(
        name = "com_github_ajg_form",
        build_file_proto_mode = "disable",
        importpath = "github.com/ajg/form",
        sum = "h1:t9c7v8JUKu/XxOGBU0yjNpaMloxGEJhUkqFRq0ibGeU=",
        version = "v1.5.1",
    )
    go_repository(
        name = "com_github_alecthomas_kong",
        build_file_proto_mode = "disable",
        importpath = "github.com/alecthomas/kong",
        sum = "h1:YIjJUiR7AcmHxL87UlbPn0gyIGwl4+nYND0OQ4ojP7k=",
        version = "v0.7.0",
    )

    go_repository(
        name = "com_github_alecthomas_template",
        build_file_proto_mode = "disable",
        importpath = "github.com/alecthomas/template",
        sum = "h1:JYp7IbQjafoB+tBA3gMyHYHrpOtNuDiK/uB5uXxq5wM=",
        version = "v0.0.0-20190718012654-fb15b899a751",
    )
    go_repository(
        name = "com_github_alecthomas_units",
        build_file_proto_mode = "disable",
        importpath = "github.com/alecthomas/units",
        sum = "h1:UQZhZ2O0vMHr2cI+DC1Mbh0TJxzA3RcLoMsFw+aXw7E=",
        version = "v0.0.0-20190924025748-f65c72e2690d",
    )
    go_repository(
        name = "com_github_aleksi_pointer",
        build_file_proto_mode = "disable",
        importpath = "github.com/AlekSi/pointer",
        sum = "h1:SSDMPcXD9jSl8FPy9cRzoRaMJtm9g9ggGTxecRUbQoI=",
        version = "v1.1.0",
    )

    go_repository(
        name = "com_github_aliyun_alibaba_cloud_sdk_go",
        build_file_proto_mode = "disable",
        importpath = "github.com/aliyun/alibaba-cloud-sdk-go",
        sum = "h1:K5561f5qv64GmtQuY0OtkY0H2x1bTXBeYwwWFPzc4D8=",
        version = "v1.62.174",
    )
    go_repository(
        name = "com_github_aliyun_aliyun_log_go_sdk",
        build_file_proto_mode = "disable",
        importpath = "github.com/aliyun/aliyun-log-go-sdk",
        sum = "h1:AmP2wVKjStvEEinj4aBiZFMchg4miIcOQ1fNwEK+5bA=",
        version = "v0.1.43",
    )

    go_repository(
        name = "com_github_aliyun_aliyun_oss_go_sdk",
        build_file_proto_mode = "disable",
        importpath = "github.com/aliyun/aliyun-oss-go-sdk",
        sum = "h1:KXeJoM1wo9I/6xPTyt6qCxoSZnmASiAjlrr0dyTUKt8=",
        version = "v2.2.6+incompatible",
    )
    go_repository(
        name = "com_github_andybalholm_brotli",
        build_file_proto_mode = "disable",
        importpath = "github.com/andybalholm/brotli",
        sum = "h1:V7DdXeJtZscaqfNuAdSRuRFzuiKlHSC/Zh3zl9qY3JY=",
        version = "v1.0.4",
    )
    go_repository(
        name = "com_github_antihax_optional",
        build_file_proto_mode = "disable",
        importpath = "github.com/antihax/optional",
        sum = "h1:xK2lYat7ZLaVVcIuj82J8kIro4V6kDe0AUDFboUCwcg=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_apache_thrift",
        build_file_proto_mode = "disable",
        importpath = "github.com/apache/thrift",
        sum = "h1:5hryIiq9gtn+MiLVn0wP37kb/uTeRZgN08WoCsAhIhI=",
        version = "v0.13.0",
    )
    go_repository(
        name = "com_github_apolloconfig_agollo_v4",
        build_file_proto_mode = "disable",
        importpath = "github.com/apolloconfig/agollo/v4",
        sum = "h1:AarIBtpDE76jIevHUOzzgT9eLB5HPvIQ8oL+OpScOPU=",
        version = "v4.3.0",
    )

    go_repository(
        name = "com_github_apparentlymart_go_dump",
        build_file_proto_mode = "disable",
        importpath = "github.com/apparentlymart/go-dump",
        sum = "h1:ZSTrOEhiM5J5RFxEaFvMZVEAM1KvT1YzbEOwB2EAGjA=",
        version = "v0.0.0-20180507223929-23540a00eaa3",
    )
    go_repository(
        name = "com_github_apparentlymart_go_textseg_v13",
        build_file_proto_mode = "disable",
        importpath = "github.com/apparentlymart/go-textseg/v13",
        sum = "h1:Y+KvPE1NYz0xl601PVImeQfFyEy6iT90AvPUL1NNfNw=",
        version = "v13.0.0",
    )
    go_repository(
        name = "com_github_armon_circbuf",
        build_file_proto_mode = "disable",
        importpath = "github.com/armon/circbuf",
        sum = "h1:QEF07wC0T1rKkctt1RINW/+RMTVmiwxETico2l3gxJA=",
        version = "v0.0.0-20150827004946-bbbad097214e",
    )
    go_repository(
        name = "com_github_armon_go_metrics",
        build_file_proto_mode = "disable",
        importpath = "github.com/armon/go-metrics",
        sum = "h1:hR91U9KYmb6bLBYLQjyM+3j+rcd/UhE+G78SFnF8gJA=",
        version = "v0.4.1",
    )
    go_repository(
        name = "com_github_armon_go_radix",
        build_file_proto_mode = "disable",
        importpath = "github.com/armon/go-radix",
        sum = "h1:F4z6KzEeeQIMeLFa97iZU6vupzoecKdU5TX24SNppXI=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_armon_go_socks5",
        build_file_proto_mode = "disable",
        importpath = "github.com/armon/go-socks5",
        sum = "h1:0CwZNZbxp69SHPdPJAN/hZIm0C4OItdklCFmMRWYpio=",
        version = "v0.0.0-20160902184237-e75332964ef5",
    )

    go_repository(
        name = "com_github_aryann_difflib",
        build_file_proto_mode = "disable",
        importpath = "github.com/aryann/difflib",
        sum = "h1:pv34s756C4pEXnjgPfGYgdhg/ZdajGhyOvzx8k+23nw=",
        version = "v0.0.0-20170710044230-e206f873d14a",
    )
    go_repository(
        name = "com_github_asaskevich_govalidator",
        build_file_proto_mode = "disable",
        importpath = "github.com/asaskevich/govalidator",
        sum = "h1:idn718Q4B6AGu/h5Sxe66HYVdqdGu2l9Iebqhi/AEoA=",
        version = "v0.0.0-20190424111038-f61b66f89f4a",
    )

    go_repository(
        name = "com_github_aws_aws_lambda_go",
        build_file_proto_mode = "disable",
        importpath = "github.com/aws/aws-lambda-go",
        sum = "h1:SuCy7H3NLyp+1Mrfp+m80jcbi9KYWAs9/BXwppwRDzY=",
        version = "v1.13.3",
    )
    go_repository(
        name = "com_github_aws_aws_sdk_go",
        build_file_proto_mode = "disable",
        importpath = "github.com/aws/aws-sdk-go",
        sum = "h1:sscPpn/Ns3i0F4HPEWAVcwdIRaZZCuL7llJ2/60yPIk=",
        version = "v1.34.28",
    )
    go_repository(
        name = "com_github_aws_aws_sdk_go_v2",
        build_file_proto_mode = "disable",
        importpath = "github.com/aws/aws-sdk-go-v2",
        sum = "h1:wyC6p9Yfq6V2y98wfDsj6OnNQa4w2BLGCLIxzNhwOGY=",
        version = "v1.17.4",
    )
    go_repository(
        name = "com_github_aws_aws_sdk_go_v2_aws_protocol_eventstream",
        build_file_proto_mode = "disable",
        importpath = "github.com/aws/aws-sdk-go-v2/aws/protocol/eventstream",
        sum = "h1:dK82zF6kkPeCo8J1e+tGx4JdvDIQzj7ygIoLg8WMuGs=",
        version = "v1.4.10",
    )
    go_repository(
        name = "com_github_aws_aws_sdk_go_v2_config",
        build_file_proto_mode = "disable",
        importpath = "github.com/aws/aws-sdk-go-v2/config",
        sum = "h1:fKs/I4wccmfrNRO9rdrbMO1NgLxct6H9rNMiPdBxHWw=",
        version = "v1.18.12",
    )
    go_repository(
        name = "com_github_aws_aws_sdk_go_v2_credentials",
        build_file_proto_mode = "disable",
        importpath = "github.com/aws/aws-sdk-go-v2/credentials",
        sum = "h1:zw1KAc1kl00NYd3ofVmFrb09qnYlSQMeh+fmlQRAihI=",
        version = "v1.13.13",
    )
    go_repository(
        name = "com_github_aws_aws_sdk_go_v2_feature_ec2_imds",
        build_file_proto_mode = "disable",
        importpath = "github.com/aws/aws-sdk-go-v2/feature/ec2/imds",
        sum = "h1:3aMfcTmoXtTZnaT86QlVaYh+BRMbvrrmZwIQ5jWqCZQ=",
        version = "v1.12.22",
    )
    go_repository(
        name = "com_github_aws_aws_sdk_go_v2_internal_configsources",
        build_file_proto_mode = "disable",
        importpath = "github.com/aws/aws-sdk-go-v2/internal/configsources",
        sum = "h1:r+XwaCLpIvCKjBIYy/HVZujQS9tsz5ohHG3ZIe0wKoE=",
        version = "v1.1.28",
    )
    go_repository(
        name = "com_github_aws_aws_sdk_go_v2_internal_endpoints_v2",
        build_file_proto_mode = "disable",
        importpath = "github.com/aws/aws-sdk-go-v2/internal/endpoints/v2",
        sum = "h1:7AwGYXDdqRQYsluvKFmWoqpcOQJ4bH634SkYf3FNj/A=",
        version = "v2.4.22",
    )
    go_repository(
        name = "com_github_aws_aws_sdk_go_v2_internal_ini",
        build_file_proto_mode = "disable",
        importpath = "github.com/aws/aws-sdk-go-v2/internal/ini",
        sum = "h1:J4xhFd6zHhdF9jPP0FQJ6WknzBboGMBNjKOv4iTuw4A=",
        version = "v1.3.29",
    )
    go_repository(
        name = "com_github_aws_aws_sdk_go_v2_internal_v4a",
        build_file_proto_mode = "disable",
        importpath = "github.com/aws/aws-sdk-go-v2/internal/v4a",
        sum = "h1:YIvKIfPXQVp0EhXUV644kmQo6cQPPSRmC44A1HSoJeg=",
        version = "v1.0.20",
    )
    go_repository(
        name = "com_github_aws_aws_sdk_go_v2_service_internal_accept_encoding",
        build_file_proto_mode = "disable",
        importpath = "github.com/aws/aws-sdk-go-v2/service/internal/accept-encoding",
        sum = "h1:y2+VQzC6Zh2ojtV2LoC0MNwHWc6qXv/j2vrQtlftkdA=",
        version = "v1.9.11",
    )
    go_repository(
        name = "com_github_aws_aws_sdk_go_v2_service_internal_checksum",
        build_file_proto_mode = "disable",
        importpath = "github.com/aws/aws-sdk-go-v2/service/internal/checksum",
        sum = "h1:c5+bNdV8E4fIPteWx4HZSkqI07oY9exbfQ7JH7Yx4PI=",
        version = "v1.1.23",
    )
    go_repository(
        name = "com_github_aws_aws_sdk_go_v2_service_internal_presigned_url",
        build_file_proto_mode = "disable",
        importpath = "github.com/aws/aws-sdk-go-v2/service/internal/presigned-url",
        sum = "h1:LjFQf8hFuMO22HkV5VWGLBvmCLBCLPivUAmpdpnp4Vs=",
        version = "v1.9.22",
    )
    go_repository(
        name = "com_github_aws_aws_sdk_go_v2_service_internal_s3shared",
        build_file_proto_mode = "disable",
        importpath = "github.com/aws/aws-sdk-go-v2/service/internal/s3shared",
        sum = "h1:ISLJ2BKXe4zzyZ7mp5ewKECiw0U7KpLgS3S6OxY9Cm0=",
        version = "v1.13.22",
    )
    go_repository(
        name = "com_github_aws_aws_sdk_go_v2_service_s3",
        build_file_proto_mode = "disable",
        importpath = "github.com/aws/aws-sdk-go-v2/service/s3",
        sum = "h1:PVieHTwugdlHedlxLpYLQsOZAq736RScuEb/m4zhzc4=",
        version = "v1.30.3",
    )
    go_repository(
        name = "com_github_aws_aws_sdk_go_v2_service_sso",
        build_file_proto_mode = "disable",
        importpath = "github.com/aws/aws-sdk-go-v2/service/sso",
        sum = "h1:EN102fWY7hI5u/2FPheTrwwMHkSXfl49RYkeEnJsrCU=",
        version = "v1.12.2",
    )
    go_repository(
        name = "com_github_aws_aws_sdk_go_v2_service_ssooidc",
        build_file_proto_mode = "disable",
        importpath = "github.com/aws/aws-sdk-go-v2/service/ssooidc",
        sum = "h1:f1lmlce7r13CX1BPyPqt9oh/H+uqOWc9367lDoGGwNQ=",
        version = "v1.14.2",
    )
    go_repository(
        name = "com_github_aws_aws_sdk_go_v2_service_sts",
        build_file_proto_mode = "disable",
        importpath = "github.com/aws/aws-sdk-go-v2/service/sts",
        sum = "h1:s49mSnsBZEXjfGBkRfmK+nPqzT7Lt3+t2SmAKNyHblw=",
        version = "v1.18.3",
    )
    go_repository(
        name = "com_github_aws_smithy_go",
        build_file_proto_mode = "disable",
        importpath = "github.com/aws/smithy-go",
        sum = "h1:hgz0X/DX0dGqTYpGALqXJoRKRj5oQ7150i5FdTePzO8=",
        version = "v1.13.5",
    )
    go_repository(
        name = "com_github_azure_go_autorest",
        build_file_proto_mode = "disable",
        importpath = "github.com/Azure/go-autorest",
        sum = "h1:V5VMDjClD3GiElqLWO7mz2MxNAK/vTfRHdAubSIPRgs=",
        version = "v14.2.0+incompatible",
    )
    go_repository(
        name = "com_github_azure_go_autorest_autorest",
        build_file_proto_mode = "disable",
        importpath = "github.com/Azure/go-autorest/autorest",
        sum = "h1:90Y4srNYrwOtAgVo3ndrQkTYn6kf1Eg/AjTFJ8Is2aM=",
        version = "v0.11.18",
    )
    go_repository(
        name = "com_github_azure_go_autorest_autorest_adal",
        build_file_proto_mode = "disable",
        importpath = "github.com/Azure/go-autorest/autorest/adal",
        sum = "h1:Mp5hbtOePIzM8pJVRa3YLrWWmZtoxRXqUEzCfJt3+/Q=",
        version = "v0.9.13",
    )
    go_repository(
        name = "com_github_azure_go_autorest_autorest_date",
        build_file_proto_mode = "disable",
        importpath = "github.com/Azure/go-autorest/autorest/date",
        sum = "h1:7gUk1U5M/CQbp9WoqinNzJar+8KY+LPI6wiWrP/myHw=",
        version = "v0.3.0",
    )
    go_repository(
        name = "com_github_azure_go_autorest_autorest_mocks",
        build_file_proto_mode = "disable",
        importpath = "github.com/Azure/go-autorest/autorest/mocks",
        sum = "h1:K0laFcLE6VLTOwNgSxaGbUcLPuGXlNkbVvq4cW4nIHk=",
        version = "v0.4.1",
    )
    go_repository(
        name = "com_github_azure_go_autorest_logger",
        build_file_proto_mode = "disable",
        importpath = "github.com/Azure/go-autorest/logger",
        sum = "h1:IG7i4p/mDa2Ce4TRyAO8IHnVhAVF3RFU+ZtXWSmf4Tg=",
        version = "v0.2.1",
    )
    go_repository(
        name = "com_github_azure_go_autorest_tracing",
        build_file_proto_mode = "disable",
        importpath = "github.com/Azure/go-autorest/tracing",
        sum = "h1:TYi4+3m5t6K48TGI9AUdb+IzbnSxvnvUMfuitfgcfuo=",
        version = "v0.6.0",
    )

    go_repository(
        name = "com_github_benbjohnson_clock",
        build_file_proto_mode = "disable",
        importpath = "github.com/benbjohnson/clock",
        sum = "h1:Q92kusRqC1XV2MjkWETPvjJVqKetz1OzxZB7mHJLju8=",
        version = "v1.1.0",
    )
    go_repository(
        name = "com_github_beorn7_perks",
        build_file_proto_mode = "disable",
        importpath = "github.com/beorn7/perks",
        sum = "h1:VlbKKnNfV8bJzeqoa4cOKqO6bYr3WgKZxO8Z16+hsOM=",
        version = "v1.0.1",
    )
    go_repository(
        name = "com_github_bgentry_speakeasy",
        build_file_proto_mode = "disable",
        importpath = "github.com/bgentry/speakeasy",
        sum = "h1:ByYyxL9InA1OWqxJqqp2A5pYHUrCiAL6K3J+LKSsQkY=",
        version = "v0.1.0",
    )
    go_repository(
        name = "com_github_bketelsen_crypt",
        build_file_proto_mode = "disable",
        importpath = "github.com/bketelsen/crypt",
        sum = "h1:+0HFd5KSZ/mm3JmhmrDukiId5iR6w4+BdFtfSy4yWIc=",
        version = "v0.0.3-0.20200106085610-5cbc8cc4026c",
    )
    go_repository(
        name = "com_github_bmizerany_assert",
        build_file_proto_mode = "disable",
        importpath = "github.com/bmizerany/assert",
        sum = "h1:DDGfHa7BWjL4YnC6+E63dPcxHo2sUxDIu8g3QgEJdRY=",
        version = "v0.0.0-20160611221934-b7ed37b82869",
    )
    go_repository(
        name = "com_github_buger_jsonparser",
        build_file_proto_mode = "disable",
        importpath = "github.com/buger/jsonparser",
        sum = "h1:2PnMjfWD7wBILjqQbt530v576A/cAbQvEW9gGIpYMUs=",
        version = "v1.1.1",
    )

    go_repository(
        name = "com_github_burntsushi_toml",
        build_file_proto_mode = "disable",
        importpath = "github.com/BurntSushi/toml",
        sum = "h1:WXkYYl6Yr3qBf1K79EBnL4mak0OimBfB0XUf9Vl28OQ=",
        version = "v0.3.1",
    )
    go_repository(
        name = "com_github_burntsushi_xgb",
        build_file_proto_mode = "disable",
        importpath = "github.com/BurntSushi/xgb",
        sum = "h1:1BDTz0u9nC3//pOCMdNH+CiXJVYJh5UQNCOBG7jbELc=",
        version = "v0.0.0-20160522181843-27f122750802",
    )
    go_repository(
        name = "com_github_casbin_casbin_v2",
        build_file_proto_mode = "disable",
        importpath = "github.com/casbin/casbin/v2",
        sum = "h1:+xd5gRceCk+IGzIJgUSZQtypZuOzZDHgHvU0FlmBCVA=",
        version = "v2.63.0",
    )
    go_repository(
        name = "com_github_casbin_ent_adapter",
        build_file_proto_mode = "disable",
        importpath = "github.com/casbin/ent-adapter",
        sum = "h1:VGpIYdkxKVTCeaxBxhbhD8dIg1a1Unxkrj4jh0GhvFk=",
        version = "v0.2.3",
    )
    go_repository(
        name = "com_github_cenkalti_backoff",
        build_file_proto_mode = "disable",
        importpath = "github.com/cenkalti/backoff",
        sum = "h1:tNowT99t7UNflLxfYYSlKYsBpXdEet03Pg2g16Swow4=",
        version = "v2.2.1+incompatible",
    )
    go_repository(
        name = "com_github_cenkalti_backoff_v4",
        build_file_proto_mode = "disable",
        importpath = "github.com/cenkalti/backoff/v4",
        sum = "h1:G2HAfAmvm/GcKan2oOQpBXOd2tT2G57ZnZGWa1PxPBQ=",
        version = "v4.1.1",
    )

    go_repository(
        name = "com_github_census_instrumentation_opencensus_proto",
        build_file_proto_mode = "disable",
        importpath = "github.com/census-instrumentation/opencensus-proto",
        sum = "h1:iKLQ0xPNFxR/2hzXZMrBo8f1j86j5WHzznCCQxV/b8g=",
        version = "v0.4.1",
    )
    go_repository(
        name = "com_github_certifi_gocertifi",
        build_file_proto_mode = "disable",
        importpath = "github.com/certifi/gocertifi",
        sum = "h1:uH66TXeswKn5PW5zdZ39xEwfS9an067BirqA+P4QaLI=",
        version = "v0.0.0-20200922220541-2c3bb06c6054",
    )

    go_repository(
        name = "com_github_cespare_xxhash",
        build_file_proto_mode = "disable",
        importpath = "github.com/cespare/xxhash",
        sum = "h1:a6HrQnmkObjyL+Gs60czilIUGqrzKutQD6XZog3p+ko=",
        version = "v1.1.0",
    )
    go_repository(
        name = "com_github_cespare_xxhash_v2",
        build_file_proto_mode = "disable",
        importpath = "github.com/cespare/xxhash/v2",
        sum = "h1:DC2CZ1Ep5Y4k3ZQ899DldepgrayRUGE6BBZ/cd9Cj44=",
        version = "v2.2.0",
    )
    go_repository(
        name = "com_github_chzyer_logex",
        build_file_proto_mode = "disable",
        importpath = "github.com/chzyer/logex",
        sum = "h1:Swpa1K6QvQznwJRcfTfQJmTE72DqScAa40E+fbHEXEE=",
        version = "v1.1.10",
    )
    go_repository(
        name = "com_github_chzyer_readline",
        build_file_proto_mode = "disable",
        importpath = "github.com/chzyer/readline",
        sum = "h1:fY5BOSpyZCqRo5OhCuC+XN+r/bBCmeuuJtjz+bCNIf8=",
        version = "v0.0.0-20180603132655-2972be24d48e",
    )
    go_repository(
        name = "com_github_chzyer_test",
        build_file_proto_mode = "disable",
        importpath = "github.com/chzyer/test",
        sum = "h1:q763qf9huN11kDQavWsoZXJNW3xEE4JJyHa5Q25/sd8=",
        version = "v0.0.0-20180213035817-a1ea475d72b1",
    )
    go_repository(
        name = "com_github_circonus_labs_circonus_gometrics",
        build_file_proto_mode = "disable",
        importpath = "github.com/circonus-labs/circonus-gometrics",
        sum = "h1:C29Ae4G5GtYyYMm1aztcyj/J5ckgJm2zwdDajFbx1NY=",
        version = "v2.3.1+incompatible",
    )
    go_repository(
        name = "com_github_circonus_labs_circonusllhist",
        build_file_proto_mode = "disable",
        importpath = "github.com/circonus-labs/circonusllhist",
        sum = "h1:TJH+oke8D16535+jHExHj4nQvzlZrj7ug5D7I/orNUA=",
        version = "v0.1.3",
    )
    go_repository(
        name = "com_github_clbanning_mxj",
        build_file_proto_mode = "disable",
        importpath = "github.com/clbanning/mxj",
        sum = "h1:HuhwZtbyvyOw+3Z1AowPkU87JkJUSv751ELWaiTpj8I=",
        version = "v1.8.4",
    )
    go_repository(
        name = "com_github_clbanning_x2j",
        build_file_proto_mode = "disable",
        importpath = "github.com/clbanning/x2j",
        sum = "h1:EdRZT3IeKQmfCSrgo8SZ8V3MEnskuJP0wCYNpe+aiXo=",
        version = "v0.0.0-20191024224557-825249438eec",
    )
    go_repository(
        name = "com_github_client9_misspell",
        build_file_proto_mode = "disable",
        importpath = "github.com/client9/misspell",
        sum = "h1:ta993UF76GwbvJcIo3Y68y/M3WxlpEHPWIGDkJYwzJI=",
        version = "v0.3.4",
    )
    go_repository(
        name = "com_github_cloudflare_golz4",
        build_file_proto_mode = "disable",
        importpath = "github.com/cloudflare/golz4",
        sum = "h1:F1EaeKL/ta07PY/k9Os/UFtwERei2/XzGemhpGnBKNg=",
        version = "v0.0.0-20150217214814-ef862a3cdc58",
    )

    go_repository(
        name = "com_github_cncf_udpa_go",
        build_file_proto_mode = "disable",
        importpath = "github.com/cncf/udpa/go",
        sum = "h1:QQ3GSy+MqSHxm/d8nCtnAiZdYFd45cYZPs8vOOIYKfk=",
        version = "v0.0.0-20220112060539-c52dc94e7fbe",
    )
    go_repository(
        name = "com_github_cncf_xds_go",
        build_file_proto_mode = "disable",
        importpath = "github.com/cncf/xds/go",
        sum = "h1:ACGZRIr7HsgBKHsueQ1yM4WaVaXh21ynwqsF8M8tXhA=",
        version = "v0.0.0-20230105202645-06c439db220b",
    )
    go_repository(
        name = "com_github_cockroachdb_apd",
        build_file_proto_mode = "disable",
        importpath = "github.com/cockroachdb/apd",
        sum = "h1:3LFP3629v+1aKXU5Q37mxmRxX/pIu1nijXydLShEq5I=",
        version = "v1.1.0",
    )
    go_repository(
        name = "com_github_cockroachdb_datadriven",
        build_file_proto_mode = "disable",
        importpath = "github.com/cockroachdb/datadriven",
        sum = "h1:xD/lrqdvwsc+O2bjSSi3YqY73Ke3LAiSCx49aCesA0E=",
        version = "v0.0.0-20200714090401-bf6692d28da5",
    )
    go_repository(
        name = "com_github_cockroachdb_errors",
        build_file_proto_mode = "disable",
        importpath = "github.com/cockroachdb/errors",
        sum = "h1:Lap807SXTH5tri2TivECb/4abUkMZC9zRoLarvcKDqs=",
        version = "v1.2.4",
    )
    go_repository(
        name = "com_github_cockroachdb_logtags",
        build_file_proto_mode = "disable",
        importpath = "github.com/cockroachdb/logtags",
        sum = "h1:o/kfcElHqOiXqcou5a3rIlMc7oJbMQkeLk0VQJ7zgqY=",
        version = "v0.0.0-20190617123548-eb05cc24525f",
    )

    go_repository(
        name = "com_github_codahale_hdrhistogram",
        build_file_proto_mode = "disable",
        importpath = "github.com/codahale/hdrhistogram",
        sum = "h1:qMd81Ts1T2OTKmB4acZcyKaMtRnY5Y44NuXGX2GFJ1w=",
        version = "v0.0.0-20161010025455-3a0bb77429bd",
    )
    go_repository(
        name = "com_github_coocood_freecache",
        build_file_proto_mode = "disable",
        importpath = "github.com/coocood/freecache",
        sum = "h1:oFyo4msX2c0QIKU+kuMJUwsKamJ+AKc2JJrKcMszJ5M=",
        version = "v1.0.1",
    )

    go_repository(
        name = "com_github_coreos_bbolt",
        build_file_proto_mode = "disable",
        importpath = "github.com/coreos/bbolt",
        sum = "h1:wZwiHHUieZCquLkDL0B8UhzreNWsPHooDAG3q34zk0s=",
        version = "v1.3.2",
    )
    go_repository(
        name = "com_github_coreos_etcd",
        build_file_proto_mode = "disable",
        importpath = "github.com/coreos/etcd",
        sum = "h1:8F3hqu9fGYLBifCmRCJsicFqDx/D68Rt3q1JMazcgBQ=",
        version = "v3.3.13+incompatible",
    )
    go_repository(
        name = "com_github_coreos_go_semver",
        build_file_proto_mode = "disable",
        importpath = "github.com/coreos/go-semver",
        sum = "h1:wkHLiw0WNATZnSG7epLsujiMCgPAc9xhjJ4tgnAxmfM=",
        version = "v0.3.0",
    )
    go_repository(
        name = "com_github_coreos_go_systemd",
        build_file_proto_mode = "disable",
        importpath = "github.com/coreos/go-systemd",
        sum = "h1:JOrtw2xFKzlg+cbHpyrpLDmnN1HqhBfnX7WDiW7eG2c=",
        version = "v0.0.0-20190719114852-fd7a80b32e1f",
    )
    go_repository(
        name = "com_github_coreos_go_systemd_v22",
        build_file_proto_mode = "disable",
        importpath = "github.com/coreos/go-systemd/v22",
        sum = "h1:rtAn27wIbmOGUs7RIbVgPEjb31ehTVniDwPGXyMxm5U=",
        version = "v22.3.3-0.20220203105225-a9a7ef127534",
    )

    go_repository(
        name = "com_github_coreos_pkg",
        build_file_proto_mode = "disable",
        importpath = "github.com/coreos/pkg",
        sum = "h1:lBNOc5arjvs8E5mO2tbpBpLoyyu8B6e44T7hJy6potg=",
        version = "v0.0.0-20180928190104-399ea9e2e55f",
    )
    go_repository(
        name = "com_github_cpuguy83_go_md2man_v2",
        build_file_proto_mode = "disable",
        importpath = "github.com/cpuguy83/go-md2man/v2",
        sum = "h1:EoUDS0afbrsXAZ9YQ9jdu/mZ2sXgT1/2yyNng4PGlyM=",
        version = "v2.0.0",
    )
    go_repository(
        name = "com_github_creack_pty",
        build_file_proto_mode = "disable",
        importpath = "github.com/creack/pty",
        sum = "h1:07n33Z8lZxZ2qwegKbObQohDhXDQxiMMz1NOUGYlesw=",
        version = "v1.1.11",
    )
    go_repository(
        name = "com_github_data_dog_go_sqlmock",
        build_file_proto_mode = "disable",
        importpath = "github.com/DATA-DOG/go-sqlmock",
        sum = "h1:Shsta01QNfFxHCfpW6YH2STWB0MudeXXEWMr20OEh60=",
        version = "v1.5.0",
    )
    go_repository(
        name = "com_github_datadog_datadog_go",
        build_file_proto_mode = "disable",
        importpath = "github.com/DataDog/datadog-go",
        sum = "h1:qSG2N4FghB1He/r2mFrWKCaL7dXCilEuNEeAn20fdD4=",
        version = "v3.2.0+incompatible",
    )
    go_repository(
        name = "com_github_davecgh_go_spew",
        build_file_proto_mode = "disable",
        importpath = "github.com/davecgh/go-spew",
        sum = "h1:vj9j/u1bqnvCEfJOwUhtlOARqs3+rkHYY13jYWTU97c=",
        version = "v1.1.1",
    )
    go_repository(
        name = "com_github_deckarep_golang_set",
        build_file_proto_mode = "disable",
        importpath = "github.com/deckarep/golang-set",
        sum = "h1:SCQV0S6gTtp6itiFrTqI+pfmJ4LN85S1YzhDf9rTHJQ=",
        version = "v1.7.1",
    )

    go_repository(
        name = "com_github_dgrijalva_jwt_go",
        build_file_proto_mode = "disable",
        importpath = "github.com/dgrijalva/jwt-go",
        sum = "h1:7qlOGliEKZXTDg6OTjfoBKDXWrumCAMpl/TFQ4/5kLM=",
        version = "v3.2.0+incompatible",
    )
    go_repository(
        name = "com_github_dgryski_go_rendezvous",
        build_file_proto_mode = "disable",
        importpath = "github.com/dgryski/go-rendezvous",
        sum = "h1:lO4WD4F/rVNCu3HqELle0jiPLLBs70cWOduZpkS1E78=",
        version = "v0.0.0-20200823014737-9f7001d12a5f",
    )

    go_repository(
        name = "com_github_dgryski_go_sip13",
        build_file_proto_mode = "disable",
        importpath = "github.com/dgryski/go-sip13",
        sum = "h1:RMLoZVzv4GliuWafOuPuQDKSm1SJph7uCRnnS61JAn4=",
        version = "v0.0.0-20181026042036-e10d5fee7954",
    )
    go_repository(
        name = "com_github_dlclark_regexp2",
        build_file_proto_mode = "disable",
        importpath = "github.com/dlclark/regexp2",
        sum = "h1:7lJfhqlPssTb1WQx4yvTHN0uElPEv52sbaECrAQxjAo=",
        version = "v1.7.0",
    )
    go_repository(
        name = "com_github_docker_spdystream",
        build_file_proto_mode = "disable",
        importpath = "github.com/docker/spdystream",
        sum = "h1:cenwrSVm+Z7QLSV/BsnenAOcDXdX4cMv4wP0B/5QbPg=",
        version = "v0.0.0-20160310174837-449fdfce4d96",
    )
    go_repository(
        name = "com_github_docopt_docopt_go",
        build_file_proto_mode = "disable",
        importpath = "github.com/docopt/docopt-go",
        sum = "h1:bWDMxwH3px2JBh6AyO7hdCn/PkvCZXii8TGj7sbtEbQ=",
        version = "v0.0.0-20180111231733-ee0de3bc6815",
    )

    go_repository(
        name = "com_github_dustin_go_humanize",
        build_file_proto_mode = "disable",
        importpath = "github.com/dustin/go-humanize",
        sum = "h1:GzkhY7T5VNhEkwH0PVJgjz+fX1rhBrR7pRT3mDkpeCY=",
        version = "v1.0.1",
    )
    go_repository(
        name = "com_github_eapache_go_resiliency",
        build_file_proto_mode = "disable",
        importpath = "github.com/eapache/go-resiliency",
        sum = "h1:RRL0nge+cWGlxXbUzJ7yMcq6w2XBEr19dCN6HECGaT0=",
        version = "v1.3.0",
    )
    go_repository(
        name = "com_github_eapache_go_xerial_snappy",
        build_file_proto_mode = "disable",
        importpath = "github.com/eapache/go-xerial-snappy",
        sum = "h1:8yY/I9ndfrgrXUbOGObLHKBR4Fl3nZXwM2c7OYTT8hM=",
        version = "v0.0.0-20230111030713-bf00bc1b83b6",
    )
    go_repository(
        name = "com_github_eapache_queue",
        build_file_proto_mode = "disable",
        importpath = "github.com/eapache/queue",
        sum = "h1:YOEu7KNc61ntiQlcEeUIoDTJ2o8mQznoNvUhiigpIqc=",
        version = "v1.1.0",
    )
    go_repository(
        name = "com_github_edsrzf_mmap_go",
        build_file_proto_mode = "disable",
        importpath = "github.com/edsrzf/mmap-go",
        sum = "h1:CEBF7HpRnUCSJgGUb5h1Gm7e3VkmVDrR8lvWVLtrOFw=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_elazarl_goproxy",
        build_file_proto_mode = "disable",
        importpath = "github.com/elazarl/goproxy",
        sum = "h1:yUdfgN0XgIJw7foRItutHYUIhlcKzcSf5vDpdhQAKTc=",
        version = "v0.0.0-20180725130230-947c36da3153",
    )
    go_repository(
        name = "com_github_emicklei_go_restful",
        build_file_proto_mode = "disable",
        importpath = "github.com/emicklei/go-restful",
        sum = "h1:SIvoTSbsMEwuM3dzFirLwKc4BH6VXP5CNf+G1FfJVr4=",
        version = "v2.12.0+incompatible",
    )
    go_repository(
        name = "com_github_emicklei_go_restful_v3",
        build_file_proto_mode = "disable",
        importpath = "github.com/emicklei/go-restful/v3",
        sum = "h1:XwGDlfxEnQZzuopoqxwSEllNcCOM9DhhFyhFIIGKwxE=",
        version = "v3.9.0",
    )

    go_repository(
        name = "com_github_envoyproxy_go_control_plane",
        build_file_proto_mode = "disable",
        importpath = "github.com/envoyproxy/go-control-plane",
        sum = "h1:xdCVXxEe0Y3FQith+0cj2irwZudqGYvecuLB1HtdexY=",
        version = "v0.10.3",
    )
    go_repository(
        name = "com_github_envoyproxy_protoc_gen_validate",
        build_file_proto_mode = "disable",
        importpath = "github.com/envoyproxy/protoc-gen-validate",
        sum = "h1:PS7VIOgmSVhWUEeZwTe7z7zouA22Cr590PzXKbZHOVY=",
        version = "v0.9.1",
    )
    go_repository(
        name = "com_github_evanphx_json_patch",
        build_file_proto_mode = "disable",
        importpath = "github.com/evanphx/json-patch",
        sum = "h1:4onqiflcdA9EOZ4RxV643DvftH5pOlLGNtQ5lPWQu84=",
        version = "v4.12.0+incompatible",
    )

    go_repository(
        name = "com_github_fasthttp_contrib_websocket",
        build_file_proto_mode = "disable",
        importpath = "github.com/fasthttp-contrib/websocket",
        sum = "h1:DddqAaWDpywytcG8w/qoQ5sAN8X12d3Z3koB0C3Rxsc=",
        version = "v0.0.0-20160511215533-1f3b11f56072",
    )
    go_repository(
        name = "com_github_fastly_go_utils",
        build_file_proto_mode = "disable",
        importpath = "github.com/fastly/go-utils",
        sum = "h1:Ghm4eQYC0nEPnSJdVkTrXpu9KtoVCSo1hg7mtI7G9KU=",
        version = "v0.0.0-20180712184237-d95a45783239",
    )

    go_repository(
        name = "com_github_fatih_color",
        build_file_proto_mode = "disable",
        importpath = "github.com/fatih/color",
        sum = "h1:8LOYc1KYPPmyKMuN8QV2DNRWNbLo6LZ0iLs8+mlH53w=",
        version = "v1.13.0",
    )
    go_repository(
        name = "com_github_fatih_structs",
        build_file_proto_mode = "disable",
        importpath = "github.com/fatih/structs",
        sum = "h1:Q7juDM0QtcnhCpeyLGQKyg4TOIghuNXrkL32pHAUMxo=",
        version = "v1.1.0",
    )
    go_repository(
        name = "com_github_felixge_httpsnoop",
        build_file_proto_mode = "disable",
        importpath = "github.com/felixge/httpsnoop",
        sum = "h1:s/nj+GCswXYzN5v2DpNMuMQYe+0DDwt5WVCU6CWBdXk=",
        version = "v1.0.3",
    )
    go_repository(
        name = "com_github_fluent_fluent_logger_golang",
        build_file_proto_mode = "disable",
        importpath = "github.com/fluent/fluent-logger-golang",
        sum = "h1:zUdY44CHX2oIUc7VTNZc+4m+ORuO/mldQDA7czhWXEg=",
        version = "v1.9.0",
    )
    go_repository(
        name = "com_github_form3tech_oss_jwt_go",
        build_file_proto_mode = "disable",
        importpath = "github.com/form3tech-oss/jwt-go",
        sum = "h1:7ZaBxOI7TMoYBfyA3cQHErNNyAWIKUMIwqxEtgHOs5c=",
        version = "v3.2.3+incompatible",
    )

    go_repository(
        name = "com_github_fortytw2_leaktest",
        build_file_proto_mode = "disable",
        importpath = "github.com/fortytw2/leaktest",
        sum = "h1:u8491cBMTQ8ft8aeV+adlcytMZylmA5nnwwkRZjI8vw=",
        version = "v1.3.0",
    )
    go_repository(
        name = "com_github_franela_goblin",
        build_file_proto_mode = "disable",
        importpath = "github.com/franela/goblin",
        sum = "h1:gb2Z18BhTPJPpLQWj4T+rfKHYCHxRHCtRxhKKjRidVw=",
        version = "v0.0.0-20200105215937-c9ffbefa60db",
    )
    go_repository(
        name = "com_github_franela_goreq",
        build_file_proto_mode = "disable",
        importpath = "github.com/franela/goreq",
        sum = "h1:a9ENSRDFBUPkJ5lCgVZh26+ZbGyoVJG7yb5SSzF5H54=",
        version = "v0.0.0-20171204163338-bcd34c9993f8",
    )
    go_repository(
        name = "com_github_frankban_quicktest",
        build_file_proto_mode = "disable",
        importpath = "github.com/frankban/quicktest",
        sum = "h1:FJKSZTDHjyhriyC81FLQ0LY93eSai0ZyR/ZIkd3ZUKE=",
        version = "v1.14.3",
    )
    go_repository(
        name = "com_github_fsnotify_fsnotify",
        build_file_proto_mode = "disable",
        importpath = "github.com/fsnotify/fsnotify",
        sum = "h1:n+5WquG0fcWoWp6xPWfHdbskMCQaFnG6PfBrh1Ky4HY=",
        version = "v1.6.0",
    )
    go_repository(
        name = "com_github_gavv_httpexpect",
        build_file_proto_mode = "disable",
        importpath = "github.com/gavv/httpexpect",
        sum = "h1:1X9kcRshkSKEjNJJxX9Y9mQ5BRfbxU5kORdjhlA1yX8=",
        version = "v2.0.0+incompatible",
    )
    go_repository(
        name = "com_github_getkin_kin_openapi",
        build_file_proto_mode = "disable",
        importpath = "github.com/getkin/kin-openapi",
        sum = "h1:j77zg3Ec+k+r+GA3d8hBoXpAc6KX9TbBPrwQGBIy2sY=",
        version = "v0.76.0",
    )
    go_repository(
        name = "com_github_getsentry_raven_go",
        build_file_proto_mode = "disable",
        importpath = "github.com/getsentry/raven-go",
        sum = "h1:no+xWJRb5ZI7eE8TWgIq1jLulQiIoLG0IfYxv5JYMGs=",
        version = "v0.2.0",
    )

    go_repository(
        name = "com_github_ghodss_yaml",
        build_file_proto_mode = "disable",
        importpath = "github.com/ghodss/yaml",
        sum = "h1:wQHKEahhL6wmXdzwWG11gIVCkOv05bNOh+Rxn0yngAk=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_gin_contrib_sse",
        build_file_proto_mode = "disable",
        importpath = "github.com/gin-contrib/sse",
        sum = "h1:Y/yl/+YNO8GZSjAhjMsSuLt29uWRFHdHYUb5lYOV9qE=",
        version = "v0.1.0",
    )
    go_repository(
        name = "com_github_gin_gonic_gin",
        build_file_proto_mode = "disable",
        importpath = "github.com/gin-gonic/gin",
        sum = "h1:UzKToD9/PoFj/V4rvlKqTRKnQYyz8Sc1MJlv4JHPtvY=",
        version = "v1.8.2",
    )
    go_repository(
        name = "com_github_go_bindata_go_bindata",
        build_file_proto_mode = "disable",
        importpath = "github.com/go-bindata/go-bindata",
        sum = "h1:WNHfSP1q2vuAa9vF54RrhCl4nqxCjVcXhlbsRXbGOSY=",
        version = "v1.0.1-0.20190711162640-ee3c2418e368",
    )
    go_repository(
        name = "com_github_go_chassis_cari",
        build_file_proto_mode = "disable",
        importpath = "github.com/go-chassis/cari",
        sum = "h1:cwBchwt9L8JOyO6QkzXFAsseMJ10zVSiVK8eDLD0HkA=",
        version = "v0.6.0",
    )
    go_repository(
        name = "com_github_go_chassis_foundation",
        build_file_proto_mode = "disable",
        importpath = "github.com/go-chassis/foundation",
        sum = "h1:z0xETnSxF+vRXWjoIhOdzt6rywjZ4sB++utEl4YgWEY=",
        version = "v0.4.0",
    )
    go_repository(
        name = "com_github_go_chassis_go_archaius",
        build_file_proto_mode = "disable",
        importpath = "github.com/go-chassis/go-archaius",
        sum = "h1:1FrNyzzmD6o6BIjPF8uQ4Cc+u7qYIgQTpDk8uopBqfo=",
        version = "v1.5.1",
    )
    go_repository(
        name = "com_github_go_chassis_go_chassis_v2",
        build_file_proto_mode = "disable",
        importpath = "github.com/go-chassis/go-chassis/v2",
        sum = "h1:qqFeRhox9Ov0GMfkFnr/ZBkhjhyGHy+SrMyDNbXl8co=",
        version = "v2.3.0",
    )
    go_repository(
        name = "com_github_go_chassis_go_restful_swagger20",
        build_file_proto_mode = "disable",
        importpath = "github.com/go-chassis/go-restful-swagger20",
        sum = "h1:kWfeLwMwJZVkXP1zNyFpkmR41UZ55UTcOptTteXhvEs=",
        version = "v1.0.3",
    )
    go_repository(
        name = "com_github_go_chassis_kie_client",
        build_file_proto_mode = "disable",
        importpath = "github.com/go-chassis/kie-client",
        sum = "h1:3vZbK1AZjYiOOXrRUdrcTNKvM1MP2qk73UyRLxaKWbI=",
        version = "v0.0.0-20201210060018-938c7680a9ab",
    )
    go_repository(
        name = "com_github_go_chassis_openlog",
        build_file_proto_mode = "disable",
        importpath = "github.com/go-chassis/openlog",
        sum = "h1:XqIOvZ8YPJ9o9lLtLBskQNNWolK5kC6a4Sv7r4s9sZ4=",
        version = "v1.1.3",
    )
    go_repository(
        name = "com_github_go_chassis_sc_client",
        build_file_proto_mode = "disable",
        importpath = "github.com/go-chassis/sc-client",
        sum = "h1:c2LSdbFMuZ3RcbDu7So//kFCzjDCkzdQ0CNKhm8Dy7I=",
        version = "v0.7.0",
    )
    go_repository(
        name = "com_github_go_chassis_seclog",
        build_file_proto_mode = "disable",
        importpath = "github.com/go-chassis/seclog",
        sum = "h1:ofjF+GpDd7YFdxa6xk13brZfVSKMdZoQVPoCWjFHGUk=",
        version = "v1.3.0",
    )
    go_repository(
        name = "com_github_go_errors_errors",
        build_file_proto_mode = "disable",
        importpath = "github.com/go-errors/errors",
        sum = "h1:LUHzmkK3GUKUrL/1gfBUxAHzcev3apQlezX/+O7ma6w=",
        version = "v1.0.1",
    )
    go_repository(
        name = "com_github_go_faster_errors",
        build_file_proto_mode = "disable",
        importpath = "github.com/go-faster/errors",
        sum = "h1:nNIPOBkprlKzkThvS/0YaX8Zs9KewLCOSFQS5BU06FI=",
        version = "v0.6.1",
    )
    go_repository(
        name = "com_github_go_faster_jx",
        build_file_proto_mode = "disable",
        importpath = "github.com/go-faster/jx",
        sum = "h1:Y9cDVPguBRX4WbZFJlEVHwA1/YDmWWOyHtzyfxaByk0=",
        version = "v0.40.0",
    )
    go_repository(
        name = "com_github_go_faster_yamlx",
        build_file_proto_mode = "disable",
        importpath = "github.com/go-faster/yamlx",
        sum = "h1:00RQkZopoLDF1SgBDJVHuN6epTOK7T0TkN427vbvEBk=",
        version = "v0.4.1",
    )

    go_repository(
        name = "com_github_go_gl_glfw",
        build_file_proto_mode = "disable",
        importpath = "github.com/go-gl/glfw",
        sum = "h1:QbL/5oDUmRBzO9/Z7Seo6zf912W/a6Sr4Eu0G/3Jho0=",
        version = "v0.0.0-20190409004039-e6da0acd62b1",
    )
    go_repository(
        name = "com_github_go_gl_glfw_v3_3_glfw",
        build_file_proto_mode = "disable",
        importpath = "github.com/go-gl/glfw/v3.3/glfw",
        sum = "h1:WtGNWLvXpe6ZudgnXrq0barxBImvnnJoMEhXAzcbM0I=",
        version = "v0.0.0-20200222043503-6f7a984d4dc4",
    )
    go_repository(
        name = "com_github_go_kit_kit",
        build_file_proto_mode = "disable",
        importpath = "github.com/go-kit/kit",
        sum = "h1:dXFJfIHVvUcpSgDOV+Ne6t7jXri8Tfv2uOLHUZ2XNuo=",
        version = "v0.10.0",
    )
    go_repository(
        name = "com_github_go_kit_log",
        build_file_proto_mode = "disable",
        importpath = "github.com/go-kit/log",
        sum = "h1:MRVx0/zhvdseW+Gza6N9rVzU/IVzaeE1SFI4raAhmBU=",
        version = "v0.2.1",
    )
    go_repository(
        name = "com_github_go_kratos_aegis",
        build_file_proto_mode = "disable",
        importpath = "github.com/go-kratos/aegis",
        sum = "h1:hXw/iO51ofO0hSRBijbmGNidthfAo9Oc/PZSl/lRCxk=",
        version = "v0.1.3",
    )
    go_repository(
        name = "com_github_go_kratos_consul",
        build_file_proto_mode = "disable",
        importpath = "github.com/go-kratos/consul",
        sum = "h1:vUWyuwwczEojXee4DJG5ae+KSqZM3qWjn1sbrOixmKY=",
        version = "v0.1.5",
    )
    go_repository(
        name = "com_github_go_kratos_grpc_gateway_v2",
        build_file_proto_mode = "disable",
        importpath = "github.com/go-kratos/grpc-gateway/v2",
        sum = "h1:jPqlxMJEoi8Yv4WIAhQNKczdKjADox8WKrTp3bJNjFM=",
        version = "v2.5.1-0.20210811062259-c92d36e434b1",
    )
    go_repository(
        name = "com_github_go_kratos_kratos_contrib_config_apollo_v2",
        build_file_proto_mode = "disable",
        importpath = "github.com/go-kratos/kratos/contrib/config/apollo/v2",
        sum = "h1:PewuFVoDh3+3POgqYvC/YQP4PIwRMYvHLFAUQQKxTlI=",
        version = "v2.0.0-20230213033822-0a076443cba1",
    )
    go_repository(
        name = "com_github_go_kratos_kratos_contrib_config_consul_v2",
        build_file_proto_mode = "disable",
        importpath = "github.com/go-kratos/kratos/contrib/config/consul/v2",
        sum = "h1:zSAZc11yb6/HfSFZ+2k/OjucclAMpowmeDEV/JgLqRQ=",
        version = "v2.0.0-20221220065744-a017ab09576f",
    )
    go_repository(
        name = "com_github_go_kratos_kratos_contrib_config_etcd_v2",
        build_file_proto_mode = "disable",
        importpath = "github.com/go-kratos/kratos/contrib/config/etcd/v2",
        sum = "h1:0Ux7G1Q3ao1/UWloTMbkGzWv7WWa9OIWEQ8TS8oUqdY=",
        version = "v2.0.0-20230213033822-0a076443cba1",
    )
    go_repository(
        name = "com_github_go_kratos_kratos_contrib_config_kubernetes_v2",
        build_file_proto_mode = "disable",
        importpath = "github.com/go-kratos/kratos/contrib/config/kubernetes/v2",
        sum = "h1:EKFd4Lq/YJ0s4R0fhKQ2yiP9G1Z89yycAevr0CsWVpw=",
        version = "v2.0.0-20230213033822-0a076443cba1",
    )
    go_repository(
        name = "com_github_go_kratos_kratos_contrib_config_nacos_v2",
        build_file_proto_mode = "disable",
        importpath = "github.com/go-kratos/kratos/contrib/config/nacos/v2",
        sum = "h1:duvTGXqQICvmvws5QhMLk42aESy/1euVLBW+fPtkv54=",
        version = "v2.0.0-20221220065744-a017ab09576f",
    )
    go_repository(
        name = "com_github_go_kratos_kratos_contrib_config_polaris_v2",
        build_file_proto_mode = "disable",
        importpath = "github.com/go-kratos/kratos/contrib/config/polaris/v2",
        sum = "h1:WeIHfkJ7/wxyDtUCqOGhwqmxWBK5wAWE1MUa49rFuew=",
        version = "v2.0.0-20230213033822-0a076443cba1",
    )
    go_repository(
        name = "com_github_go_kratos_kratos_contrib_log_aliyun_v2",
        build_file_proto_mode = "disable",
        importpath = "github.com/go-kratos/kratos/contrib/log/aliyun/v2",
        sum = "h1:3ov4OAyjApAawfkq60zxAXfLIoHhT7b+z0V6GMQyb/s=",
        version = "v2.0.0-20230227131608-c65f823c38de",
    )
    go_repository(
        name = "com_github_go_kratos_kratos_contrib_log_fluent_v2",
        build_file_proto_mode = "disable",
        importpath = "github.com/go-kratos/kratos/contrib/log/fluent/v2",
        sum = "h1:0x6imrXSZM9E7Xas26/S3+ILvlPzTPdqByy/kjKUpjA=",
        version = "v2.0.0-20230227131608-c65f823c38de",
    )
    go_repository(
        name = "com_github_go_kratos_kratos_contrib_log_logrus_v2",
        build_file_proto_mode = "disable",
        importpath = "github.com/go-kratos/kratos/contrib/log/logrus/v2",
        sum = "h1:d3aLHbWi9kh1FaQDcEzT5/D4VonJyZ9bXWtvSEdryiI=",
        version = "v2.0.0-20230227131608-c65f823c38de",
    )
    go_repository(
        name = "com_github_go_kratos_kratos_contrib_log_tencent_v2",
        build_file_proto_mode = "disable",
        importpath = "github.com/go-kratos/kratos/contrib/log/tencent/v2",
        sum = "h1:XqMdY2Lz/pHKBue+7BXFq+f/kMMNdEOHuzrs73gf77Y=",
        version = "v2.0.0-20230227131608-c65f823c38de",
    )
    go_repository(
        name = "com_github_go_kratos_kratos_contrib_log_zap_v2",
        build_file_proto_mode = "disable",
        importpath = "github.com/go-kratos/kratos/contrib/log/zap/v2",
        sum = "h1:oEwOrZ26wS81f25uMup4c/AgGVlRpOaSi0ADnXkQpwo=",
        version = "v2.0.0-20230227131608-c65f823c38de",
    )

    go_repository(
        name = "com_github_go_kratos_kratos_contrib_metrics_prometheus_v2",
        build_file_proto_mode = "disable",
        importpath = "github.com/go-kratos/kratos/contrib/metrics/prometheus/v2",
        sum = "h1:oXzl/WyflN1hkltaOIt7jcfr8i8J+ioI8/bs/5F0hgM=",
        version = "v2.0.0-20221220065744-a017ab09576f",
    )
    go_repository(
        name = "com_github_go_kratos_kratos_contrib_registry_consul_v2",
        build_file_proto_mode = "disable",
        importpath = "github.com/go-kratos/kratos/contrib/registry/consul/v2",
        sum = "h1:PrXrZa+hMsf7xtwECCRVW2dyHZt8+35UHSz1KsFcTjU=",
        version = "v2.0.0-20230213033822-0a076443cba1",
    )
    go_repository(
        name = "com_github_go_kratos_kratos_contrib_registry_etcd_v2",
        build_file_proto_mode = "disable",
        importpath = "github.com/go-kratos/kratos/contrib/registry/etcd/v2",
        sum = "h1:H3IP87PvMvuIVamy/qjlbDWSRfLdA5r9kYp1YXNXXb0=",
        version = "v2.0.0-20230213033822-0a076443cba1",
    )
    go_repository(
        name = "com_github_go_kratos_kratos_contrib_registry_eureka_v2",
        build_file_proto_mode = "disable",
        importpath = "github.com/go-kratos/kratos/contrib/registry/eureka/v2",
        sum = "h1:fXqjYlCNerBbRoBRqj3qnh9FjaJITbZ/9wFdUcVieCo=",
        version = "v2.0.0-20230213033822-0a076443cba1",
    )
    go_repository(
        name = "com_github_go_kratos_kratos_contrib_registry_kubernetes_v2",
        build_file_proto_mode = "disable",
        importpath = "github.com/go-kratos/kratos/contrib/registry/kubernetes/v2",
        sum = "h1:vmNDcu2sY3+AH/NafuCpv3F/MfTrOZJm9aj5cqbqwVk=",
        version = "v2.0.0-20230213033822-0a076443cba1",
    )
    go_repository(
        name = "com_github_go_kratos_kratos_contrib_registry_nacos_v2",
        build_file_proto_mode = "disable",
        importpath = "github.com/go-kratos/kratos/contrib/registry/nacos/v2",
        sum = "h1:DifE/XtTb3RH4HxSaZy3AfDSCCF0x2g8Zzy21JDRPKI=",
        version = "v2.0.0-20230213033822-0a076443cba1",
    )
    go_repository(
        name = "com_github_go_kratos_kratos_contrib_registry_polaris_v2",
        build_file_proto_mode = "disable",
        importpath = "github.com/go-kratos/kratos/contrib/registry/polaris/v2",
        sum = "h1:6+ckESgwy24MFwi7aDgmfGGzJNcFDmzgREOQbdjRakg=",
        version = "v2.0.0-20230213033822-0a076443cba1",
    )
    go_repository(
        name = "com_github_go_kratos_kratos_contrib_registry_servicecomb_v2",
        build_file_proto_mode = "disable",
        importpath = "github.com/go-kratos/kratos/contrib/registry/servicecomb/v2",
        sum = "h1:urNDg+t7eed4SE9HuKcxZhPzPXB36LpCAjS0hq721/8=",
        version = "v2.0.0-20230213033822-0a076443cba1",
    )
    go_repository(
        name = "com_github_go_kratos_kratos_contrib_registry_zookeeper_v2",
        build_file_proto_mode = "disable",
        importpath = "github.com/go-kratos/kratos/contrib/registry/zookeeper/v2",
        sum = "h1:+kwzcWhb9pKKxYdvy8WLK31Q5Zw+1agnPkPNCd/OSpE=",
        version = "v2.0.0-20230213033822-0a076443cba1",
    )

    go_repository(
        name = "com_github_go_kratos_kratos_v2",
        build_file_proto_mode = "disable",
        importpath = "github.com/go-kratos/kratos/v2",
        sum = "h1:tZeZXqEQ6nUPxE82TENDCzJ5y/295oo/vFg6P4G2tQY=",
        version = "v2.5.4",
    )
    go_repository(
        name = "com_github_go_kratos_swagger_api",
        build_file_proto_mode = "disable",
        importpath = "github.com/go-kratos/swagger-api",
        sum = "h1:zBagw2tWA+icYrRPFqM3jFL60nO743ZiIvuiI6CpV6E=",
        version = "v1.0.1",
    )
    go_repository(
        name = "com_github_go_logfmt_logfmt",
        build_file_proto_mode = "disable",
        importpath = "github.com/go-logfmt/logfmt",
        sum = "h1:otpy5pqBCBZ1ng9RQ0dPu4PN7ba75Y/aA+UpowDyNVA=",
        version = "v0.5.1",
    )
    go_repository(
        name = "com_github_go_logr_logr",
        build_file_proto_mode = "disable",
        importpath = "github.com/go-logr/logr",
        sum = "h1:2DntVwHkVopvECVRSlL5PSo9eG+cAkDCuckLubN+rq0=",
        version = "v1.2.3",
    )
    go_repository(
        name = "com_github_go_logr_stdr",
        build_file_proto_mode = "disable",
        importpath = "github.com/go-logr/stdr",
        sum = "h1:hSWxHoqTgW2S2qGc0LTAI563KZ5YKYRhT3MFKZMbjag=",
        version = "v1.2.2",
    )
    go_repository(
        name = "com_github_go_oauth2_oauth2_v4",
        build_file_proto_mode = "disable",
        importpath = "github.com/go-oauth2/oauth2/v4",
        sum = "h1:CuZhD3lhGuI6aNLyUbRHXsgG2RwGRBOuCBfd4WQKqBQ=",
        version = "v4.5.2",
    )
    go_repository(
        name = "com_github_go_ole_go_ole",
        build_file_proto_mode = "disable",
        importpath = "github.com/go-ole/go-ole",
        sum = "h1:/Fpf6oFPoeFik9ty7siob0G6Ke8QvQEuVcuChpwXzpY=",
        version = "v1.2.6",
    )
    go_repository(
        name = "com_github_go_openapi_inflect",
        build_file_proto_mode = "disable",
        importpath = "github.com/go-openapi/inflect",
        sum = "h1:9jCH9scKIbHeV9m12SmPilScz6krDxKRasNNSNPXu/4=",
        version = "v0.19.0",
    )
    go_repository(
        name = "com_github_go_openapi_jsonpointer",
        build_file_proto_mode = "disable",
        importpath = "github.com/go-openapi/jsonpointer",
        sum = "h1:gZr+CIYByUqjcgeLXnQu2gHYQC9o73G2XUeOFYEICuY=",
        version = "v0.19.5",
    )
    go_repository(
        name = "com_github_go_openapi_jsonreference",
        build_file_proto_mode = "disable",
        importpath = "github.com/go-openapi/jsonreference",
        sum = "h1:MYlu0sBgChmCfJxxUKZ8g1cPWFOB37YSZqewK7OKeyA=",
        version = "v0.20.0",
    )
    go_repository(
        name = "com_github_go_openapi_spec",
        build_file_proto_mode = "disable",
        importpath = "github.com/go-openapi/spec",
        sum = "h1:C1JKChikHGpXwT5UQDFaryIpDtyyGL/CR6C2kB7F1oc=",
        version = "v0.0.0-20160808142527-6aced65f8501",
    )
    go_repository(
        name = "com_github_go_openapi_swag",
        build_file_proto_mode = "disable",
        importpath = "github.com/go-openapi/swag",
        sum = "h1:gm3vOOXfiuw5i9p5N9xJvfjvuofpyvLA9Wr6QfK5Fng=",
        version = "v0.19.14",
    )

    go_repository(
        name = "com_github_go_playground_assert_v2",
        build_file_proto_mode = "disable",
        importpath = "github.com/go-playground/assert/v2",
        sum = "h1:JvknZsQTYeFEAhQwI4qEt9cyV5ONwRHC+lYKSsYSR8s=",
        version = "v2.2.0",
    )
    go_repository(
        name = "com_github_go_playground_form_v4",
        build_file_proto_mode = "disable",
        importpath = "github.com/go-playground/form/v4",
        sum = "h1:N1wh+Goz61e6w66vo8vJkQt+uwZSoLz50kZPJWR8eic=",
        version = "v4.2.0",
    )
    go_repository(
        name = "com_github_go_playground_locales",
        build_file_proto_mode = "disable",
        importpath = "github.com/go-playground/locales",
        sum = "h1:u50s323jtVGugKlcYeyzC0etD1HifMjqmJqb8WugfUU=",
        version = "v0.14.0",
    )
    go_repository(
        name = "com_github_go_playground_universal_translator",
        build_file_proto_mode = "disable",
        importpath = "github.com/go-playground/universal-translator",
        sum = "h1:82dyy6p4OuJq4/CByFNOn/jYrnRPArHwAcmLoJZxyho=",
        version = "v0.18.0",
    )
    go_repository(
        name = "com_github_go_playground_validator",
        build_file_proto_mode = "disable",
        importpath = "github.com/go-playground/validator",
        sum = "h1:UA72EPEogEnq76ehGdEDp4Mit+3FDh548oRqwVgNsHA=",
        version = "v9.31.0+incompatible",
    )

    go_repository(
        name = "com_github_go_playground_validator_v10",
        build_file_proto_mode = "disable",
        importpath = "github.com/go-playground/validator/v10",
        sum = "h1:prmOlTVv+YjZjmRmNSF3VmspqJIxJWXmqUsHwfTRRkQ=",
        version = "v10.11.1",
    )
    go_repository(
        name = "com_github_go_redis_redis",
        build_file_proto_mode = "disable",
        importpath = "github.com/go-redis/redis",
        sum = "h1:K0pv1D7EQUjfyoMql+r/jZqCLizCGKFlFgcHWWmHQjg=",
        version = "v6.15.9+incompatible",
    )
    go_repository(
        name = "com_github_go_redis_redis_extra_rediscmd_v8",
        build_file_proto_mode = "disable",
        importpath = "github.com/go-redis/redis/extra/rediscmd/v8",
        sum = "h1:ftG8tp8SG81xyuL2woNEx5t2RZ8mOJuC2+tumi+/NR8=",
        version = "v8.11.5",
    )
    go_repository(
        name = "com_github_go_redis_redis_extra_redisotel_v8",
        build_file_proto_mode = "disable",
        importpath = "github.com/go-redis/redis/extra/redisotel/v8",
        sum = "h1:BqyYJgvdSr2S/6O2l7zmCj26ocUTxDLgagsGIRfkS+Q=",
        version = "v8.11.5",
    )
    go_repository(
        name = "com_github_go_redis_redis_v8",
        build_file_proto_mode = "disable",
        importpath = "github.com/go-redis/redis/v8",
        sum = "h1:AcZZR7igkdvfVmQTPnu9WE37LRrO/YrBH5zWyjDC0oI=",
        version = "v8.11.5",
    )

    go_repository(
        name = "com_github_go_resty_resty_v2",
        build_file_proto_mode = "disable",
        importpath = "github.com/go-resty/resty/v2",
        sum = "h1:me+K9p3uhSmXtrBZ4k9jcEAfJmuC8IivWHwaLZwPrFY=",
        version = "v2.7.0",
    )
    go_repository(
        name = "com_github_go_session_session",
        build_file_proto_mode = "disable",
        importpath = "github.com/go-session/session",
        sum = "h1:yStchEObKg4nk2F7JGE7KoFIrA/1Y078peagMWcrncg=",
        version = "v3.1.2+incompatible",
    )
    go_repository(
        name = "com_github_go_sql_driver_mysql",
        build_file_proto_mode = "disable",
        importpath = "github.com/go-sql-driver/mysql",
        sum = "h1:ueSltNNllEqE3qcWBTD0iQd3IpL/6U+mJxLkazJ7YPc=",
        version = "v1.7.0",
    )
    go_repository(
        name = "com_github_go_stack_stack",
        build_file_proto_mode = "disable",
        importpath = "github.com/go-stack/stack",
        sum = "h1:5SgMzNM5HxrEjV0ww2lTmX6E2Izsfxas4+YHWRs3Lsk=",
        version = "v1.8.0",
    )
    go_repository(
        name = "com_github_go_task_slim_sprig",
        build_file_proto_mode = "disable",
        importpath = "github.com/go-task/slim-sprig",
        sum = "h1:p104kn46Q8WdvHunIJ9dAyjPVtrBPhSr3KT2yUst43I=",
        version = "v0.0.0-20210107165309-348f09dbbbc0",
    )

    go_repository(
        name = "com_github_go_test_deep",
        build_file_proto_mode = "disable",
        importpath = "github.com/go-test/deep",
        sum = "h1:ZrJSEWsXzPOxaZnFteGEfooLba+ju3FYIbOrS+rQd68=",
        version = "v1.0.3",
    )
    go_repository(
        name = "com_github_go_zookeeper_zk",
        build_file_proto_mode = "disable",
        importpath = "github.com/go-zookeeper/zk",
        sum = "h1:7M2kwOsc//9VeeFiPtf+uSJlVpU66x9Ba5+8XK7/TDg=",
        version = "v1.0.3",
    )
    go_repository(
        name = "com_github_gobuffalo_attrs",
        build_file_proto_mode = "disable",
        importpath = "github.com/gobuffalo/attrs",
        sum = "h1:hSkbZ9XSyjyBirMeqSqUrK+9HboWrweVlzRNqoBi2d4=",
        version = "v0.0.0-20190224210810-a9411de4debd",
    )
    go_repository(
        name = "com_github_gobuffalo_depgen",
        build_file_proto_mode = "disable",
        importpath = "github.com/gobuffalo/depgen",
        sum = "h1:31atYa/UW9V5q8vMJ+W6wd64OaaTHUrCUXER358zLM4=",
        version = "v0.1.0",
    )
    go_repository(
        name = "com_github_gobuffalo_envy",
        build_file_proto_mode = "disable",
        importpath = "github.com/gobuffalo/envy",
        sum = "h1:GlXgaiBkmrYMHco6t4j7SacKO4XUjvh5pwXh0f4uxXU=",
        version = "v1.7.0",
    )
    go_repository(
        name = "com_github_gobuffalo_flect",
        build_file_proto_mode = "disable",
        importpath = "github.com/gobuffalo/flect",
        sum = "h1:3GQ53z7E3o00C/yy7Ko8VXqQXoJGLkrTQCLTF1EjoXU=",
        version = "v0.1.3",
    )
    go_repository(
        name = "com_github_gobuffalo_genny",
        build_file_proto_mode = "disable",
        importpath = "github.com/gobuffalo/genny",
        sum = "h1:iQ0D6SpNXIxu52WESsD+KoQ7af2e3nCfnSBoSF/hKe0=",
        version = "v0.1.1",
    )
    go_repository(
        name = "com_github_gobuffalo_gitgen",
        build_file_proto_mode = "disable",
        importpath = "github.com/gobuffalo/gitgen",
        sum = "h1:mSVZ4vj4khv+oThUfS+SQU3UuFIZ5Zo6UNcvK8E8Mz8=",
        version = "v0.0.0-20190315122116-cc086187d211",
    )
    go_repository(
        name = "com_github_gobuffalo_gogen",
        build_file_proto_mode = "disable",
        importpath = "github.com/gobuffalo/gogen",
        sum = "h1:dLg+zb+uOyd/mKeQUYIbwbNmfRsr9hd/WtYWepmayhI=",
        version = "v0.1.1",
    )
    go_repository(
        name = "com_github_gobuffalo_logger",
        build_file_proto_mode = "disable",
        importpath = "github.com/gobuffalo/logger",
        sum = "h1:8thhT+kUJMTMy3HlX4+y9Da+BNJck+p109tqqKp7WDs=",
        version = "v0.0.0-20190315122211-86e12af44bc2",
    )
    go_repository(
        name = "com_github_gobuffalo_mapi",
        build_file_proto_mode = "disable",
        importpath = "github.com/gobuffalo/mapi",
        sum = "h1:fq9WcL1BYrm36SzK6+aAnZ8hcp+SrmnDyAxhNx8dvJk=",
        version = "v1.0.2",
    )
    go_repository(
        name = "com_github_gobuffalo_packd",
        build_file_proto_mode = "disable",
        importpath = "github.com/gobuffalo/packd",
        sum = "h1:4sGKOD8yaYJ+dek1FDkwcxCHA40M4kfKgFHx8N2kwbU=",
        version = "v0.1.0",
    )
    go_repository(
        name = "com_github_gobuffalo_packr_v2",
        build_file_proto_mode = "disable",
        importpath = "github.com/gobuffalo/packr/v2",
        sum = "h1:Ir9W9XIm9j7bhhkKE9cokvtTl1vBm62A/fene/ZCj6A=",
        version = "v2.2.0",
    )
    go_repository(
        name = "com_github_gobuffalo_syncx",
        build_file_proto_mode = "disable",
        importpath = "github.com/gobuffalo/syncx",
        sum = "h1:tpom+2CJmpzAWj5/VEHync2rJGi+epHNIeRSWjzGA+4=",
        version = "v0.0.0-20190224160051-33c29581e754",
    )

    go_repository(
        name = "com_github_goccy_go_json",
        build_file_proto_mode = "disable",
        importpath = "github.com/goccy/go-json",
        sum = "h1:mXKd9Qw4NuzShiRlOXKews24ufknHO7gx30lsDyokKA=",
        version = "v0.10.0",
    )
    go_repository(
        name = "com_github_godbus_dbus_v5",
        build_file_proto_mode = "disable",
        importpath = "github.com/godbus/dbus/v5",
        sum = "h1:9349emZab16e7zQvpmsbtjc18ykshndd8y2PG3sgJbA=",
        version = "v5.0.4",
    )

    go_repository(
        name = "com_github_gofrs_uuid",
        build_file_proto_mode = "disable",
        importpath = "github.com/gofrs/uuid",
        sum = "h1:yyYWMnhkhrKwwr8gAOcOCYxOOscHgDS9yZgBrnJfGa0=",
        version = "v4.2.0+incompatible",
    )
    go_repository(
        name = "com_github_gogo_googleapis",
        build_file_proto_mode = "disable",
        importpath = "github.com/gogo/googleapis",
        sum = "h1:kFkMAZBNAn4j7K0GiZr8cRYzejq68VbheufiV3YuyFI=",
        version = "v1.1.0",
    )
    go_repository(
        name = "com_github_gogo_protobuf",
        build_file_proto_mode = "disable",
        importpath = "github.com/gogo/protobuf",
        sum = "h1:Ov1cvc58UF3b5XjBnZv7+opcTcQFZebYjWzi34vdm4Q=",
        version = "v1.3.2",
    )
    go_repository(
        name = "com_github_goji_httpauth",
        build_file_proto_mode = "disable",
        importpath = "github.com/goji/httpauth",
        sum = "h1:lBXNCxVENCipq4D1Is42JVOP4eQjlB8TQ6H69Yx5J9Q=",
        version = "v0.0.0-20160601135302-2da839ab0f4d",
    )
    go_repository(
        name = "com_github_golang_freetype",
        build_file_proto_mode = "disable",
        importpath = "github.com/golang/freetype",
        sum = "h1:DACJavvAHhabrF08vX0COfcOBJRhZ8lUbR+ZWIs0Y5g=",
        version = "v0.0.0-20170609003504-e2365dfdc4a0",
    )
    go_repository(
        name = "com_github_golang_glog",
        build_file_proto_mode = "disable",
        importpath = "github.com/golang/glog",
        sum = "h1:nfP3RFugxnNRyKgeWd4oI1nYvXpxrx8ck8ZrcizshdQ=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_golang_groupcache",
        build_file_proto_mode = "disable",
        importpath = "github.com/golang/groupcache",
        sum = "h1:oI5xCqsCo564l8iNU+DwB5epxmsaqB+rhGL0m5jtYqE=",
        version = "v0.0.0-20210331224755-41bb18bfe9da",
    )
    go_repository(
        name = "com_github_golang_jwt_jwt",
        build_file_proto_mode = "disable",
        importpath = "github.com/golang-jwt/jwt",
        sum = "h1:IfV12K8xAKAnZqdXVzCZ+TOjboZ2keLg81eXfW3O+oY=",
        version = "v3.2.2+incompatible",
    )
    go_repository(
        name = "com_github_golang_jwt_jwt_v4",
        build_file_proto_mode = "disable",
        importpath = "github.com/golang-jwt/jwt/v4",
        sum = "h1:Hxl6lhQFj4AnOX6MLrsCb/+7tCj7DxP7VA+2rDIq5AU=",
        version = "v4.4.3",
    )
    go_repository(
        name = "com_github_golang_mock",
        build_file_proto_mode = "disable",
        importpath = "github.com/golang/mock",
        sum = "h1:ErTB+efbowRARo13NNdxyJji2egdxLGQhRaY+DUumQc=",
        version = "v1.6.0",
    )
    go_repository(
        name = "com_github_golang_protobuf",
        build_file_proto_mode = "disable",
        importpath = "github.com/golang/protobuf",
        sum = "h1:ROPKBNFfQgOUMifHyP+KYbvpjbdoFNs+aK7DXlji0Tw=",
        version = "v1.5.2",
    )
    go_repository(
        name = "com_github_golang_snappy",
        build_file_proto_mode = "disable",
        importpath = "github.com/golang/snappy",
        sum = "h1:yAGX7huGHXlcLOEtBnF4w7FQwA26wojNCwOYAEhLjQM=",
        version = "v0.0.4",
    )
    go_repository(
        name = "com_github_gonum_blas",
        build_file_proto_mode = "disable",
        importpath = "github.com/gonum/blas",
        sum = "h1:Q0Jsdxl5jbxouNs1TQYt0gxesYMU4VXRbsTlgDloZ50=",
        version = "v0.0.0-20181208220705-f22b278b28ac",
    )
    go_repository(
        name = "com_github_gonum_floats",
        build_file_proto_mode = "disable",
        importpath = "github.com/gonum/floats",
        sum = "h1:EvokxLQsaaQjcWVWSV38221VAK7qc2zhaO17bKys/18=",
        version = "v0.0.0-20181209220543-c233463c7e82",
    )
    go_repository(
        name = "com_github_gonum_integrate",
        build_file_proto_mode = "disable",
        importpath = "github.com/gonum/integrate",
        sum = "h1:GUSkTcIe1SlregbHNUKbYDhBsS8lNgYfIp4S4cToUyU=",
        version = "v0.0.0-20181209220457-a422b5c0fdf2",
    )
    go_repository(
        name = "com_github_gonum_internal",
        build_file_proto_mode = "disable",
        importpath = "github.com/gonum/internal",
        sum = "h1:8jtTdc+Nfj9AR+0soOeia9UZSvYBvETVHZrugUowJ7M=",
        version = "v0.0.0-20181124074243-f884aa714029",
    )
    go_repository(
        name = "com_github_gonum_lapack",
        build_file_proto_mode = "disable",
        importpath = "github.com/gonum/lapack",
        sum = "h1:7qnwS9+oeSiOIsiUMajT+0R7HR6hw5NegnKPmn/94oI=",
        version = "v0.0.0-20181123203213-e4cdc5a0bff9",
    )
    go_repository(
        name = "com_github_gonum_matrix",
        build_file_proto_mode = "disable",
        importpath = "github.com/gonum/matrix",
        sum = "h1:V2IgdyerlBa/MxaEFRbV5juy/C3MGdj4ePi+g6ePIp4=",
        version = "v0.0.0-20181209220409-c518dec07be9",
    )
    go_repository(
        name = "com_github_gonum_stat",
        build_file_proto_mode = "disable",
        importpath = "github.com/gonum/stat",
        sum = "h1:fbskpz/cPqWH8VqkQ7LJghFkl2KPAiIFUHrTJ2O3RGk=",
        version = "v0.0.0-20181125101827-41a0da705a5b",
    )

    go_repository(
        name = "com_github_google_btree",
        build_file_proto_mode = "disable",
        importpath = "github.com/google/btree",
        sum = "h1:gK4Kx5IaGY9CD5sPJ36FHiBJ6ZXl0kilRiiCj+jdYp4=",
        version = "v1.0.1",
    )
    go_repository(
        name = "com_github_google_gnostic",
        build_file_proto_mode = "disable",
        importpath = "github.com/google/gnostic",
        sum = "h1:FhTMOKj2VhjpouxvWJAV1TL304uMlb9zcDqkl6cEI54=",
        version = "v0.5.7-v3refs",
    )

    go_repository(
        name = "com_github_google_go_cmp",
        build_file_proto_mode = "disable",
        importpath = "github.com/google/go-cmp",
        sum = "h1:O2Tfq5qg4qc4AmwVlvv0oLiVAGB7enBSJ2x2DqQFi38=",
        version = "v0.5.9",
    )
    go_repository(
        name = "com_github_google_go_querystring",
        build_file_proto_mode = "disable",
        importpath = "github.com/google/go-querystring",
        sum = "h1:AnCroh3fv4ZBgVIf1Iwtovgjaw/GiKJo8M8yD/fhyJ8=",
        version = "v1.1.0",
    )
    go_repository(
        name = "com_github_google_gofuzz",
        build_file_proto_mode = "disable",
        importpath = "github.com/google/gofuzz",
        sum = "h1:Hsa8mG0dQ46ij8Sl2AYJDUv1oA9/d6Vk+3LG99Oe02g=",
        version = "v1.1.0",
    )
    go_repository(
        name = "com_github_google_martian",
        build_file_proto_mode = "disable",
        importpath = "github.com/google/martian",
        sum = "h1:/CP5g8u/VJHijgedC/Legn3BAbAaWPgecwXBIDzw5no=",
        version = "v2.1.0+incompatible",
    )
    go_repository(
        name = "com_github_google_martian_v3",
        build_file_proto_mode = "disable",
        importpath = "github.com/google/martian/v3",
        sum = "h1:d8MncMlErDFTwQGBK1xhv026j9kqhvw1Qv9IbWT1VLQ=",
        version = "v3.2.1",
    )
    go_repository(
        name = "com_github_google_pprof",
        build_file_proto_mode = "disable",
        importpath = "github.com/google/pprof",
        sum = "h1:K6RDEckDVWvDI9JAJYCmNdQXq6neHJOYx3V6jnqNEec=",
        version = "v0.0.0-20210720184732-4bb14d4b1be1",
    )
    go_repository(
        name = "com_github_google_renameio",
        build_file_proto_mode = "disable",
        importpath = "github.com/google/renameio",
        sum = "h1:GOZbcHa3HfsPKPlmyPyN2KEohoMXOhdMbHrvbpl2QaA=",
        version = "v0.1.0",
    )
    go_repository(
        name = "com_github_google_subcommands",
        build_file_proto_mode = "disable",
        importpath = "github.com/google/subcommands",
        sum = "h1:/eqq+otEXm5vhfBrbREPCSVQbvofip6kIz+mX5TUH7k=",
        version = "v1.0.1",
    )
    go_repository(
        name = "com_github_google_uuid",
        build_file_proto_mode = "disable",
        importpath = "github.com/google/uuid",
        sum = "h1:t6JiXgmwXMjEs8VusXIJk2BXHsn+wx8BZdTaoZ5fu7I=",
        version = "v1.3.0",
    )
    go_repository(
        name = "com_github_google_wire",
        build_file_proto_mode = "disable",
        importpath = "github.com/google/wire",
        sum = "h1:I7ELFeVBr3yfPIcc8+MWvrjk+3VjbcSzoXm3JVa+jD8=",
        version = "v0.5.0",
    )
    go_repository(
        name = "com_github_googleapis_enterprise_certificate_proxy",
        build_file_proto_mode = "disable",
        importpath = "github.com/googleapis/enterprise-certificate-proxy",
        sum = "h1:y8Yozv7SZtlU//QXbezB6QkpuE6jMD2/gfzk4AftXjs=",
        version = "v0.2.0",
    )

    go_repository(
        name = "com_github_googleapis_gax_go_v2",
        build_file_proto_mode = "disable",
        importpath = "github.com/googleapis/gax-go/v2",
        sum = "h1:SXk3ABtQYDT/OH8jAyvEOQ58mgawq5C4o/4/89qN2ZU=",
        version = "v2.6.0",
    )
    go_repository(
        name = "com_github_googleapis_gnostic",
        build_file_proto_mode = "disable",
        importpath = "github.com/googleapis/gnostic",
        sum = "h1:7XGaL1e6bYS1yIonGp9761ExpPPV1ui0SAC59Yube9k=",
        version = "v0.0.0-20170729233727-0c5108395e2d",
    )
    go_repository(
        name = "com_github_googleapis_google_cloud_go_testing",
        build_file_proto_mode = "disable",
        importpath = "github.com/googleapis/google-cloud-go-testing",
        sum = "h1:tlyzajkF3030q6M8SvmJSemC9DTHL/xaMa18b65+JM4=",
        version = "v0.0.0-20200911160855-bcd43fbb19e8",
    )
    go_repository(
        name = "com_github_gophercloud_gophercloud",
        build_file_proto_mode = "disable",
        importpath = "github.com/gophercloud/gophercloud",
        sum = "h1:P/nh25+rzXouhytV2pUHBb65fnds26Ghl8/391+sT5o=",
        version = "v0.1.0",
    )

    go_repository(
        name = "com_github_gopherjs_gopherjs",
        build_file_proto_mode = "disable",
        importpath = "github.com/gopherjs/gopherjs",
        sum = "h1:EGx4pi6eqNxGaHF6qqu48+N2wcFQ5qg5FXgOdqsJ5d8=",
        version = "v0.0.0-20181017120253-0766667cb4d1",
    )
    go_repository(
        name = "com_github_gordonklaus_ineffassign",
        build_file_proto_mode = "disable",
        importpath = "github.com/gordonklaus/ineffassign",
        sum = "h1:vc7Dmrk4JwS0ZPS6WZvWlwDflgDTA26jItmbSj83nug=",
        version = "v0.0.0-20200309095847-7953dde2c7bf",
    )

    go_repository(
        name = "com_github_gorilla_context",
        build_file_proto_mode = "disable",
        importpath = "github.com/gorilla/context",
        sum = "h1:AWwleXJkX/nhcU9bZSnZoi3h/qGYqQAGhq6zZe/aQW8=",
        version = "v1.1.1",
    )
    go_repository(
        name = "com_github_gorilla_handlers",
        build_file_proto_mode = "disable",
        importpath = "github.com/gorilla/handlers",
        sum = "h1:9lRY6j8DEeeBT10CvO9hGW0gmky0BprnvDI5vfhUHH4=",
        version = "v1.5.1",
    )
    go_repository(
        name = "com_github_gorilla_mux",
        build_file_proto_mode = "disable",
        importpath = "github.com/gorilla/mux",
        sum = "h1:i40aqfkR1h2SlN9hojwV5ZA91wcXFOvkdNIeFDP5koI=",
        version = "v1.8.0",
    )
    go_repository(
        name = "com_github_gorilla_securecookie",
        build_file_proto_mode = "disable",
        importpath = "github.com/gorilla/securecookie",
        sum = "h1:miw7JPhV+b/lAHSXz4qd/nN9jRiAFV5FwjeKyCS8BvQ=",
        version = "v1.1.1",
    )
    go_repository(
        name = "com_github_gorilla_sessions",
        build_file_proto_mode = "disable",
        importpath = "github.com/gorilla/sessions",
        sum = "h1:DHd3rPN5lE3Ts3D8rKkQ8x/0kqfeNmBAaiSi+o7FsgI=",
        version = "v1.2.1",
    )
    go_repository(
        name = "com_github_gorilla_websocket",
        build_file_proto_mode = "disable",
        importpath = "github.com/gorilla/websocket",
        sum = "h1:PPwGk2jz7EePpoHN/+ClbZu8SPxiqlu12wZP/3sWmnc=",
        version = "v1.5.0",
    )
    go_repository(
        name = "com_github_gregjones_httpcache",
        build_file_proto_mode = "disable",
        importpath = "github.com/gregjones/httpcache",
        sum = "h1:pdN6V1QBWetyv/0+wjACpqVH+eVULgEjkurDLq3goeM=",
        version = "v0.0.0-20180305231024-9cad4c3443a7",
    )

    go_repository(
        name = "com_github_grpc_ecosystem_go_grpc_middleware",
        build_file_proto_mode = "disable",
        importpath = "github.com/grpc-ecosystem/go-grpc-middleware",
        sum = "h1:+9834+KizmvFV7pXQGSXQTsaWhq2GjuNUt0aUU0YBYw=",
        version = "v1.3.0",
    )
    go_repository(
        name = "com_github_grpc_ecosystem_go_grpc_prometheus",
        build_file_proto_mode = "disable",
        importpath = "github.com/grpc-ecosystem/go-grpc-prometheus",
        sum = "h1:Ovs26xHkKqVztRpIrF/92BcuyuQ/YW4NSIpoGtfXNho=",
        version = "v1.2.0",
    )
    go_repository(
        name = "com_github_grpc_ecosystem_grpc_gateway",
        build_file_proto_mode = "disable",
        importpath = "github.com/grpc-ecosystem/grpc-gateway",
        sum = "h1:gmcG1KaJ57LophUzW0Hy8NmPhnMZb4M0+kPpLofRdBo=",
        version = "v1.16.0",
    )
    go_repository(
        name = "com_github_grpc_ecosystem_grpc_gateway_v2",
        build_file_proto_mode = "disable",
        importpath = "github.com/grpc-ecosystem/grpc-gateway/v2",
        sum = "h1:1JYBfzqrWPcCclBwxFCPAou9n+q86mfnu7NAeHfte7A=",
        version = "v2.15.0",
    )
    go_repository(
        name = "com_github_hashicorp_consul_api",
        build_file_proto_mode = "disable",
        importpath = "github.com/hashicorp/consul/api",
        sum = "h1:R7PPNzTCeN6VuQNDwwhZWJvzCtGSrNpJqfb22h3yH9g=",
        version = "v1.18.0",
    )
    go_repository(
        name = "com_github_hashicorp_consul_sdk",
        build_file_proto_mode = "disable",
        importpath = "github.com/hashicorp/consul/sdk",
        sum = "h1:lce3nFlpv8humJL8rNrrGHYSKc3q+Kxfeg3Ii1m6ZWU=",
        version = "v0.13.0",
    )
    go_repository(
        name = "com_github_hashicorp_errwrap",
        build_file_proto_mode = "disable",
        importpath = "github.com/hashicorp/errwrap",
        sum = "h1:OxrOeh75EUXMY8TBjag2fzXGZ40LB6IKw45YeGUDY2I=",
        version = "v1.1.0",
    )
    go_repository(
        name = "com_github_hashicorp_go_cleanhttp",
        build_file_proto_mode = "disable",
        importpath = "github.com/hashicorp/go-cleanhttp",
        sum = "h1:035FKYIWjmULyFRBKPs8TBQoi0x6d9G4xc9neXJWAZQ=",
        version = "v0.5.2",
    )
    go_repository(
        name = "com_github_hashicorp_go_hclog",
        build_file_proto_mode = "disable",
        importpath = "github.com/hashicorp/go-hclog",
        sum = "h1:ctuWFGrhFha8BnnzxqeRGidlEcQkDyL5u8J8t5eA11I=",
        version = "v1.4.0",
    )
    go_repository(
        name = "com_github_hashicorp_go_immutable_radix",
        build_file_proto_mode = "disable",
        importpath = "github.com/hashicorp/go-immutable-radix",
        sum = "h1:DKHmCUm2hRBK510BaiZlwvpD40f8bJFeZnpfm2KLowc=",
        version = "v1.3.1",
    )
    go_repository(
        name = "com_github_hashicorp_go_msgpack",
        build_file_proto_mode = "disable",
        importpath = "github.com/hashicorp/go-msgpack",
        sum = "h1:i9R9JSrqIz0QVLz3sz+i3YJdT7TTSLcfLLzJi9aZTuI=",
        version = "v0.5.5",
    )
    go_repository(
        name = "com_github_hashicorp_go_multierror",
        build_file_proto_mode = "disable",
        importpath = "github.com/hashicorp/go-multierror",
        sum = "h1:H5DkEtf6CXdFp0N0Em5UCwQpXMWke8IA0+lD48awMYo=",
        version = "v1.1.1",
    )
    go_repository(
        name = "com_github_hashicorp_go_net",
        build_file_proto_mode = "disable",
        importpath = "github.com/hashicorp/go.net",
        sum = "h1:sNCoNyDEvN1xa+X0baata4RdcpKwcMS6DH+xwfqPgjw=",
        version = "v0.0.1",
    )
    go_repository(
        name = "com_github_hashicorp_go_retryablehttp",
        build_file_proto_mode = "disable",
        importpath = "github.com/hashicorp/go-retryablehttp",
        sum = "h1:QlWt0KvWT0lq8MFppF9tsJGF+ynG7ztc2KIPhzRGk7s=",
        version = "v0.5.3",
    )
    go_repository(
        name = "com_github_hashicorp_go_rootcerts",
        build_file_proto_mode = "disable",
        importpath = "github.com/hashicorp/go-rootcerts",
        sum = "h1:jzhAVGtqPKbwpyCPELlgNWhE1znq+qwJtW5Oi2viEzc=",
        version = "v1.0.2",
    )
    go_repository(
        name = "com_github_hashicorp_go_sockaddr",
        build_file_proto_mode = "disable",
        importpath = "github.com/hashicorp/go-sockaddr",
        sum = "h1:ztczhD1jLxIRjVejw8gFomI1BQZOe2WoVOu0SyteCQc=",
        version = "v1.0.2",
    )
    go_repository(
        name = "com_github_hashicorp_go_syslog",
        build_file_proto_mode = "disable",
        importpath = "github.com/hashicorp/go-syslog",
        sum = "h1:KaodqZuhUoZereWVIYmpUgZysurB1kBLX2j0MwMrUAE=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_hashicorp_go_uuid",
        build_file_proto_mode = "disable",
        importpath = "github.com/hashicorp/go-uuid",
        sum = "h1:2gKiV6YVmrJ1i2CKKa9obLvRieoRGviZFL26PcT/Co8=",
        version = "v1.0.3",
    )
    go_repository(
        name = "com_github_hashicorp_go_version",
        build_file_proto_mode = "disable",
        importpath = "github.com/hashicorp/go-version",
        sum = "h1:zEfKbn2+PDgroKdiOzqiE8rsmLqU2uwi5PB5pBJ3TkI=",
        version = "v1.2.1",
    )
    go_repository(
        name = "com_github_hashicorp_golang_lru",
        build_file_proto_mode = "disable",
        importpath = "github.com/hashicorp/golang-lru",
        sum = "h1:YDjusn29QI/Das2iO9M0BHnIbxPeyuCHsjMW+lJfyTc=",
        version = "v0.5.4",
    )
    go_repository(
        name = "com_github_hashicorp_hcl",
        build_file_proto_mode = "disable",
        importpath = "github.com/hashicorp/hcl",
        sum = "h1:0Anlzjpi4vEasTeNFn2mLJgTSwt0+6sfsiTG8qcWGx4=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_hashicorp_hcl_v2",
        build_file_proto_mode = "disable",
        importpath = "github.com/hashicorp/hcl/v2",
        sum = "h1:CPDXO6+uORPjKflkWCCwoWc9uRp+zSIPcCQ+BrxV7m8=",
        version = "v2.15.0",
    )
    go_repository(
        name = "com_github_hashicorp_logutils",
        build_file_proto_mode = "disable",
        importpath = "github.com/hashicorp/logutils",
        sum = "h1:dLEQVugN8vlakKOUE3ihGLTZJRB4j+M2cdTm/ORI65Y=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_hashicorp_mdns",
        build_file_proto_mode = "disable",
        importpath = "github.com/hashicorp/mdns",
        sum = "h1:sY0CMhFmjIPDMlTB+HfymFHCaYLhgifZ0QhjaYKD/UQ=",
        version = "v1.0.4",
    )
    go_repository(
        name = "com_github_hashicorp_memberlist",
        build_file_proto_mode = "disable",
        importpath = "github.com/hashicorp/memberlist",
        sum = "h1:EtYPN8DpAURiapus508I4n9CzHs2W+8NZGbmmR/prTM=",
        version = "v0.5.0",
    )
    go_repository(
        name = "com_github_hashicorp_serf",
        build_file_proto_mode = "disable",
        importpath = "github.com/hashicorp/serf",
        sum = "h1:Z1H2J60yRKvfDYAOZLd2MU0ND4AH/WDz7xYHDWQsIPY=",
        version = "v0.10.1",
    )
    go_repository(
        name = "com_github_hpcloud_tail",
        build_file_proto_mode = "disable",
        importpath = "github.com/hpcloud/tail",
        sum = "h1:nfCOvKYfkgYP8hkirhJocXT2+zOD8yUNjXaWfTlyFKI=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_hudl_fargo",
        build_file_proto_mode = "disable",
        importpath = "github.com/hudl/fargo",
        sum = "h1:0U6+BtN6LhaYuTnIJq4Wyq5cpn6O2kWrxAtcqBmYY6w=",
        version = "v1.3.0",
    )
    go_repository(
        name = "com_github_iancoleman_strcase",
        build_file_proto_mode = "disable",
        importpath = "github.com/iancoleman/strcase",
        sum = "h1:05I4QRnGpI0m37iZQRuskXh+w77mr6Z41lwQzuHLwW0=",
        version = "v0.2.0",
    )
    go_repository(
        name = "com_github_ianlancetaylor_demangle",
        build_file_proto_mode = "disable",
        importpath = "github.com/ianlancetaylor/demangle",
        sum = "h1:mV02weKRL81bEnm8A0HT1/CAelMQDBuQIfLw8n+d6xI=",
        version = "v0.0.0-20200824232613-28f6c0f3b639",
    )
    go_repository(
        name = "com_github_imdario_mergo",
        build_file_proto_mode = "disable",
        importpath = "github.com/imdario/mergo",
        sum = "h1:lFzP57bqS/wsqKssCGmtLAb8A0wKjLGrve2q3PPVcBk=",
        version = "v0.3.13",
    )
    go_repository(
        name = "com_github_imkira_go_interpol",
        build_file_proto_mode = "disable",
        importpath = "github.com/imkira/go-interpol",
        sum = "h1:KIiKr0VSG2CUW1hl1jpiyuzuJeKUUpC8iM1AIE7N1Vk=",
        version = "v1.1.0",
    )
    go_repository(
        name = "com_github_inconshreveable_mousetrap",
        build_file_proto_mode = "disable",
        importpath = "github.com/inconshreveable/mousetrap",
        sum = "h1:U3uMjPSQEBMNp1lFxmllqCPM6P5u/Xq7Pgzkat/bFNc=",
        version = "v1.0.1",
    )
    go_repository(
        name = "com_github_influxdata_influxdb1_client",
        build_file_proto_mode = "disable",
        importpath = "github.com/influxdata/influxdb1-client",
        sum = "h1:/WZQPMZNsjZ7IlCpsLGdQBINg5bxKQ1K1sh6awxLtkA=",
        version = "v0.0.0-20191209144304-8bf82d3c094d",
    )
    go_repository(
        name = "com_github_jackc_chunkreader",
        build_file_proto_mode = "disable",
        importpath = "github.com/jackc/chunkreader",
        sum = "h1:4s39bBR8ByfqH+DKm8rQA3E1LHZWB9XWcrz8fqaZbe0=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_jackc_chunkreader_v2",
        build_file_proto_mode = "disable",
        importpath = "github.com/jackc/chunkreader/v2",
        sum = "h1:i+RDz65UE+mmpjTfyz0MoVTnzeYxroil2G82ki7MGG8=",
        version = "v2.0.1",
    )
    go_repository(
        name = "com_github_jackc_pgconn",
        build_file_proto_mode = "disable",
        importpath = "github.com/jackc/pgconn",
        sum = "h1:3L1XMNV2Zvca/8BYhzcRFS70Lr0WlDg16Di6SFGAbys=",
        version = "v1.13.0",
    )
    go_repository(
        name = "com_github_jackc_pgio",
        build_file_proto_mode = "disable",
        importpath = "github.com/jackc/pgio",
        sum = "h1:g12B9UwVnzGhueNavwioyEEpAmqMe1E/BN9ES+8ovkE=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_jackc_pgmock",
        build_file_proto_mode = "disable",
        importpath = "github.com/jackc/pgmock",
        sum = "h1:DadwsjnMwFjfWc9y5Wi/+Zz7xoE5ALHsRQlOctkOiHc=",
        version = "v0.0.0-20210724152146-4ad1a8207f65",
    )
    go_repository(
        name = "com_github_jackc_pgpassfile",
        build_file_proto_mode = "disable",
        importpath = "github.com/jackc/pgpassfile",
        sum = "h1:/6Hmqy13Ss2zCq62VdNG8tM1wchn8zjSGOBJ6icpsIM=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_jackc_pgproto3",
        build_file_proto_mode = "disable",
        importpath = "github.com/jackc/pgproto3",
        sum = "h1:FYYE4yRw+AgI8wXIinMlNjBbp/UitDJwfj5LqqewP1A=",
        version = "v1.1.0",
    )
    go_repository(
        name = "com_github_jackc_pgproto3_v2",
        build_file_proto_mode = "disable",
        importpath = "github.com/jackc/pgproto3/v2",
        sum = "h1:nwj7qwf0S+Q7ISFfBndqeLwSwxs+4DPsbRFjECT1Y4Y=",
        version = "v2.3.1",
    )
    go_repository(
        name = "com_github_jackc_pgservicefile",
        build_file_proto_mode = "disable",
        importpath = "github.com/jackc/pgservicefile",
        sum = "h1:bbPeKD0xmW/Y25WS6cokEszi5g+S0QxI/d45PkRi7Nk=",
        version = "v0.0.0-20221227161230-091c0ba34f0a",
    )
    go_repository(
        name = "com_github_jackc_pgtype",
        build_file_proto_mode = "disable",
        importpath = "github.com/jackc/pgtype",
        sum = "h1:XkIc7A+1BmZD19bB2NxrtjJweHxQ9agqvM+9URc68Cg=",
        version = "v1.13.0",
    )
    go_repository(
        name = "com_github_jackc_pgx_v4",
        build_file_proto_mode = "disable",
        importpath = "github.com/jackc/pgx/v4",
        sum = "h1:0Ut0rpeKwvIVbMQ1KbMBU4h6wxehBI535LK6Flheh8E=",
        version = "v4.17.2",
    )
    go_repository(
        name = "com_github_jackc_puddle",
        build_file_proto_mode = "disable",
        importpath = "github.com/jackc/puddle",
        sum = "h1:eHK/5clGOatcjX3oWGBO/MpxpbHzSwud5EWTSCI+MX0=",
        version = "v1.3.0",
    )
    go_repository(
        name = "com_github_jcmturner_aescts_v2",
        build_file_proto_mode = "disable",
        importpath = "github.com/jcmturner/aescts/v2",
        sum = "h1:9YKLH6ey7H4eDBXW8khjYslgyqG2xZikXP0EQFKrle8=",
        version = "v2.0.0",
    )
    go_repository(
        name = "com_github_jcmturner_dnsutils_v2",
        build_file_proto_mode = "disable",
        importpath = "github.com/jcmturner/dnsutils/v2",
        sum = "h1:lltnkeZGL0wILNvrNiVCR6Ro5PGU/SeBvVO/8c/iPbo=",
        version = "v2.0.0",
    )
    go_repository(
        name = "com_github_jcmturner_gofork",
        build_file_proto_mode = "disable",
        importpath = "github.com/jcmturner/gofork",
        sum = "h1:QH0l3hzAU1tfT3rZCnW5zXl+orbkNMMRGJfdJjHVETg=",
        version = "v1.7.6",
    )
    go_repository(
        name = "com_github_jcmturner_goidentity_v6",
        build_file_proto_mode = "disable",
        importpath = "github.com/jcmturner/goidentity/v6",
        sum = "h1:VKnZd2oEIMorCTsFBnJWbExfNN7yZr3EhJAxwOkZg6o=",
        version = "v6.0.1",
    )
    go_repository(
        name = "com_github_jcmturner_gokrb5_v8",
        build_file_proto_mode = "disable",
        importpath = "github.com/jcmturner/gokrb5/v8",
        sum = "h1:iTonLeSJOn7MVUtyMT+arAn5AKAPrkilzhGw8wE/Tq8=",
        version = "v8.4.3",
    )
    go_repository(
        name = "com_github_jcmturner_rpc_v2",
        build_file_proto_mode = "disable",
        importpath = "github.com/jcmturner/rpc/v2",
        sum = "h1:7FXXj8Ti1IaVFpSAziCZWNzbNuZmnvw/i6CqLNdWfZY=",
        version = "v2.0.3",
    )
    go_repository(
        name = "com_github_jehiah_go_strftime",
        build_file_proto_mode = "disable",
        importpath = "github.com/jehiah/go-strftime",
        sum = "h1:IPJ3dvxmJ4uczJe5YQdrYB16oTJlGSC/OyZDqUk9xX4=",
        version = "v0.0.0-20171201141054-1d33003b3869",
    )

    go_repository(
        name = "com_github_jessevdk_go_flags",
        build_file_proto_mode = "disable",
        importpath = "github.com/jessevdk/go-flags",
        sum = "h1:1jKYvbxEjfUl0fmqTCOfonvskHHXMjBySTLW4y9LFvc=",
        version = "v1.5.0",
    )
    go_repository(
        name = "com_github_jhump_protoreflect",
        build_file_proto_mode = "disable",
        importpath = "github.com/jhump/protoreflect",
        sum = "h1:iH+UZfsbRE6vpyZH7asAjTPWJf7RJbpZ9j/N3lDlKs0=",
        version = "v1.10.1",
    )

    go_repository(
        name = "com_github_jinzhu_copier",
        build_file_proto_mode = "disable",
        importpath = "github.com/jinzhu/copier",
        sum = "h1:GlvfUwHk62RokgqVNvYsku0TATCF7bAHVwEXoBh3iJg=",
        version = "v0.3.5",
    )
    go_repository(
        name = "com_github_jmespath_go_jmespath",
        build_file_proto_mode = "disable",
        importpath = "github.com/jmespath/go-jmespath",
        sum = "h1:BEgLn5cpjn8UN1mAw4NjwDrS35OdebyEtFe+9YPoQUg=",
        version = "v0.4.0",
    )
    go_repository(
        name = "com_github_jmespath_go_jmespath_internal_testify",
        build_file_proto_mode = "disable",
        importpath = "github.com/jmespath/go-jmespath/internal/testify",
        sum = "h1:shLQSRRSCCPj3f2gpwzGwWFoC7ycTf1rcQZHOlsJ6N8=",
        version = "v1.5.1",
    )
    go_repository(
        name = "com_github_joho_godotenv",
        build_file_proto_mode = "disable",
        importpath = "github.com/joho/godotenv",
        sum = "h1:Zjp+RcGpHhGlrMbJzXTrZZPrWj+1vfm90La1wgB6Bhc=",
        version = "v1.3.0",
    )

    go_repository(
        name = "com_github_jonboulle_clockwork",
        build_file_proto_mode = "disable",
        importpath = "github.com/jonboulle/clockwork",
        sum = "h1:UOGuzwb1PwsrDAObMuhUnj0p5ULPj8V/xJ7Kx9qUBdQ=",
        version = "v0.2.2",
    )
    go_repository(
        name = "com_github_josharian_intern",
        build_file_proto_mode = "disable",
        importpath = "github.com/josharian/intern",
        sum = "h1:vlS4z54oSdjm0bgjRigI+G1HpF+tI+9rE5LLzOg8HmY=",
        version = "v1.0.0",
    )

    go_repository(
        name = "com_github_jpillora_backoff",
        build_file_proto_mode = "disable",
        importpath = "github.com/jpillora/backoff",
        sum = "h1:uvFg412JmmHBHw7iwprIxkPMI+sGQ4kzOWsMeHnm2EA=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_json_iterator_go",
        build_file_proto_mode = "disable",
        importpath = "github.com/json-iterator/go",
        sum = "h1:PV8peI4a0ysnczrg+LtxykD8LfKY9ML6u2jnxaEnrnM=",
        version = "v1.1.12",
    )
    go_repository(
        name = "com_github_jstemmer_go_junit_report",
        build_file_proto_mode = "disable",
        importpath = "github.com/jstemmer/go-junit-report",
        sum = "h1:6QPYqodiu3GuPL+7mfx+NwDdp2eTkp9IfEUpgAwUN0o=",
        version = "v0.9.1",
    )
    go_repository(
        name = "com_github_jtolds_gls",
        build_file_proto_mode = "disable",
        importpath = "github.com/jtolds/gls",
        sum = "h1:xdiiI2gbIgH/gLH7ADydsJ1uDOEzR8yvV7C0MuV77Wo=",
        version = "v4.20.0+incompatible",
    )
    go_repository(
        name = "com_github_julienschmidt_httprouter",
        build_file_proto_mode = "disable",
        importpath = "github.com/julienschmidt/httprouter",
        sum = "h1:U0609e9tgbseu3rBINet9P48AI/D3oJs4dN7jwJOQ1U=",
        version = "v1.3.0",
    )
    go_repository(
        name = "com_github_k0kubun_colorstring",
        build_file_proto_mode = "disable",
        importpath = "github.com/k0kubun/colorstring",
        sum = "h1:uC1QfSlInpQF+M0ao65imhwqKnz3Q2z/d8PWZRMQvDM=",
        version = "v0.0.0-20150214042306-9440f1994b88",
    )
    go_repository(
        name = "com_github_karlseguin_ccache_v2",
        build_file_proto_mode = "disable",
        importpath = "github.com/karlseguin/ccache/v2",
        sum = "h1:lT38cE//uyf6KcFok0rlgXtGFBWxkI6h/qg4tbFyDnA=",
        version = "v2.0.8",
    )
    go_repository(
        name = "com_github_karlseguin_expect",
        build_file_proto_mode = "disable",
        importpath = "github.com/karlseguin/expect",
        sum = "h1:vJ0Snvo+SLMY72r5J4sEfkuE7AFbixEP2qRbEcum/wA=",
        version = "v1.0.2-0.20190806010014-778a5f0c6003",
    )
    go_repository(
        name = "com_github_karrick_godirwalk",
        build_file_proto_mode = "disable",
        importpath = "github.com/karrick/godirwalk",
        sum = "h1:lOpSw2vJP0y5eLBW906QwKsUK/fe/QDyoqM5rnnuPDY=",
        version = "v1.10.3",
    )

    go_repository(
        name = "com_github_kisielk_errcheck",
        build_file_proto_mode = "disable",
        importpath = "github.com/kisielk/errcheck",
        sum = "h1:e8esj/e4R+SAOwFwN+n3zr0nYeCyeweozKfO23MvHzY=",
        version = "v1.5.0",
    )
    go_repository(
        name = "com_github_kisielk_gotool",
        build_file_proto_mode = "disable",
        importpath = "github.com/kisielk/gotool",
        sum = "h1:AV2c/EiW3KqPNT9ZKl07ehoAGi4C5/01Cfbblndcapg=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_klauspost_compress",
        build_file_proto_mode = "disable",
        importpath = "github.com/klauspost/compress",
        sum = "h1:EF27CXIuDsYJ6mmvtBRlEuB2UVOqHG1tAXgZ7yIO+lw=",
        version = "v1.15.15",
    )
    go_repository(
        name = "com_github_klauspost_cpuid_v2",
        build_file_proto_mode = "disable",
        importpath = "github.com/klauspost/cpuid/v2",
        sum = "h1:sxCkb+qR91z4vsqw4vGGZlDgPz3G7gjaLyK3V8y70BU=",
        version = "v2.2.3",
    )

    go_repository(
        name = "com_github_knetic_govaluate",
        build_file_proto_mode = "disable",
        importpath = "github.com/Knetic/govaluate",
        sum = "h1:1G1pk05UrOh0NlF1oeaaix1x8XzrfjIDK47TY0Zehcw=",
        version = "v3.0.1-0.20171022003610-9aa49832a739+incompatible",
    )
    go_repository(
        name = "com_github_konsorten_go_windows_terminal_sequences",
        build_file_proto_mode = "disable",
        importpath = "github.com/konsorten/go-windows-terminal-sequences",
        sum = "h1:CE8S1cTafDpPvMhIxNJKvHsGVBgn1xWYf1NbHQhywc8=",
        version = "v1.0.3",
    )
    go_repository(
        name = "com_github_kr_fs",
        build_file_proto_mode = "disable",
        importpath = "github.com/kr/fs",
        sum = "h1:Jskdu9ieNAYnjxsi0LbQp1ulIKZV1LAFgK1tWhpZgl8=",
        version = "v0.1.0",
    )

    go_repository(
        name = "com_github_kr_logfmt",
        build_file_proto_mode = "disable",
        importpath = "github.com/kr/logfmt",
        sum = "h1:T+h1c/A9Gawja4Y9mFVWj2vyii2bbUNDw3kt9VxK2EY=",
        version = "v0.0.0-20140226030751-b84e30acd515",
    )
    go_repository(
        name = "com_github_kr_pretty",
        build_file_proto_mode = "disable",
        importpath = "github.com/kr/pretty",
        sum = "h1:WgNl7dwNpEZ6jJ9k1snq4pZsg7DOEN8hP9Xw0Tsjwk0=",
        version = "v0.3.0",
    )
    go_repository(
        name = "com_github_kr_pty",
        build_file_proto_mode = "disable",
        importpath = "github.com/kr/pty",
        sum = "h1:AkaSdXYQOWeaO3neb8EM634ahkXXe3jYbVh/F9lq+GI=",
        version = "v1.1.8",
    )
    go_repository(
        name = "com_github_kr_text",
        build_file_proto_mode = "disable",
        importpath = "github.com/kr/text",
        sum = "h1:5Nx0Ya0ZqY2ygV366QzturHI13Jq95ApcVaJBhpS+AY=",
        version = "v0.2.0",
    )
    go_repository(
        name = "com_github_kylelemons_godebug",
        build_file_proto_mode = "disable",
        importpath = "github.com/kylelemons/godebug",
        sum = "h1:MtvEpTB6LX3vkb4ax0b5D2DHbNAUsen0Gx5wZoq3lV4=",
        version = "v0.0.0-20170820004349-d65d576e9348",
    )
    go_repository(
        name = "com_github_leodido_go_urn",
        build_file_proto_mode = "disable",
        importpath = "github.com/leodido/go-urn",
        sum = "h1:BqpAaACuzVSgi/VLzGZIobT2z4v53pjosyNd9Yv6n/w=",
        version = "v1.2.1",
    )
    go_repository(
        name = "com_github_lestrrat_go_envload",
        build_file_proto_mode = "disable",
        importpath = "github.com/lestrrat/go-envload",
        sum = "h1:0iQektZGS248WXmGIYOwRXSQhD4qn3icjMpuxwO7qlo=",
        version = "v0.0.0-20180220120943-6ed08b54a570",
    )
    go_repository(
        name = "com_github_lestrrat_go_file_rotatelogs",
        build_file_proto_mode = "disable",
        importpath = "github.com/lestrrat/go-file-rotatelogs",
        sum = "h1:sgUSP4zdTUZYZgAGGtN5Lxk92rK+JUFOwf+FT99EEI4=",
        version = "v0.0.0-20180223000712-d3151e2a480f",
    )
    go_repository(
        name = "com_github_lestrrat_go_strftime",
        build_file_proto_mode = "disable",
        importpath = "github.com/lestrrat/go-strftime",
        sum = "h1:Bvq8AziQ5jFF4BHGAEDSqwPW1NJS3XshxbRCxtjFAZc=",
        version = "v0.0.0-20180220042222-ba3bf9c1d042",
    )

    go_repository(
        name = "com_github_lib_pq",
        build_file_proto_mode = "disable",
        importpath = "github.com/lib/pq",
        sum = "h1:p7ZhMD+KsSRozJr34udlUrhboJwWAgCg34+/ZZNvZZw=",
        version = "v1.10.7",
    )
    go_repository(
        name = "com_github_lightstep_lightstep_tracer_common_golang_gogo",
        build_file_proto_mode = "disable",
        importpath = "github.com/lightstep/lightstep-tracer-common/golang/gogo",
        sum = "h1:143Bb8f8DuGWck/xpNUOckBVYfFbBTnLevfRZ1aVVqo=",
        version = "v0.0.0-20190605223551-bc2310a04743",
    )
    go_repository(
        name = "com_github_lightstep_lightstep_tracer_go",
        build_file_proto_mode = "disable",
        importpath = "github.com/lightstep/lightstep-tracer-go",
        sum = "h1:vi1F1IQ8N7hNWytK9DpJsUfQhGuNSc19z330K6vl4zk=",
        version = "v0.18.1",
    )
    go_repository(
        name = "com_github_little_cui_etcdadpt",
        build_file_proto_mode = "disable",
        importpath = "github.com/little-cui/etcdadpt",
        sum = "h1:EBXPBxddZXTgWvGsIdAqqG6JCu1TouPNUhVVj9swt/s=",
        version = "v0.3.2",
    )
    go_repository(
        name = "com_github_logrusorgru_aurora_v3",
        build_file_proto_mode = "disable",
        importpath = "github.com/logrusorgru/aurora/v3",
        sum = "h1:R6zcoZZbvVcGMvDCKo45A9U/lzYyzl5NfYIvznmDfE4=",
        version = "v3.0.0",
    )

    go_repository(
        name = "com_github_lufia_plan9stats",
        build_file_proto_mode = "disable",
        importpath = "github.com/lufia/plan9stats",
        sum = "h1:VtwQ41oftZwlMnOEbMWQtSEUgU64U4s+GHk7hZK+jtY=",
        version = "v0.0.0-20220913051719-115f729f3c8c",
    )
    go_repository(
        name = "com_github_lyft_protoc_gen_star",
        build_file_proto_mode = "disable",
        importpath = "github.com/lyft/protoc-gen-star",
        sum = "h1:erE0rdztuaDq3bpGifD95wfoPrSZc95nGA6tbiNYh6M=",
        version = "v0.6.1",
    )
    go_repository(
        name = "com_github_lyft_protoc_gen_validate",
        build_file_proto_mode = "disable",
        importpath = "github.com/lyft/protoc-gen-validate",
        sum = "h1:KNt/RhmQTOLr7Aj8PsJ7mTronaFyx80mRTT9qF261dA=",
        version = "v0.0.13",
    )
    go_repository(
        name = "com_github_magiconair_properties",
        build_file_proto_mode = "disable",
        importpath = "github.com/magiconair/properties",
        sum = "h1:IeQXZAiQcpL9mgcAe1Nu6cX9LLw6ExEHKjN0VQdvPDY=",
        version = "v1.8.7",
    )
    go_repository(
        name = "com_github_mailru_easyjson",
        build_file_proto_mode = "disable",
        importpath = "github.com/mailru/easyjson",
        sum = "h1:8yTIVnZgCoiM1TgqoeTl+LfU5Jg6/xL3QhGQnimLYnA=",
        version = "v0.7.6",
    )
    go_repository(
        name = "com_github_markbates_oncer",
        build_file_proto_mode = "disable",
        importpath = "github.com/markbates/oncer",
        sum = "h1:JgVTCPf0uBVcUSWpyXmGpgOc62nK5HWUBKAGc3Qqa5k=",
        version = "v0.0.0-20181203154359-bf2de49a0be2",
    )
    go_repository(
        name = "com_github_markbates_safe",
        build_file_proto_mode = "disable",
        importpath = "github.com/markbates/safe",
        sum = "h1:yjZkbvRM6IzKj9tlu/zMJLS0n/V351OZWRnF3QfaUxI=",
        version = "v1.0.1",
    )

    go_repository(
        name = "com_github_masterminds_semver_v3",
        build_file_proto_mode = "disable",
        importpath = "github.com/Masterminds/semver/v3",
        sum = "h1:hLg3sBzpNErnxhQtUy/mmLR2I9foDujNK030IGemrRc=",
        version = "v3.1.1",
    )
    go_repository(
        name = "com_github_mattn_go_colorable",
        build_file_proto_mode = "disable",
        importpath = "github.com/mattn/go-colorable",
        sum = "h1:fFA4WZxdEF4tXPZVKMLwD8oUnCTTo08duU7wxecdEvA=",
        version = "v0.1.13",
    )
    go_repository(
        name = "com_github_mattn_go_isatty",
        build_file_proto_mode = "disable",
        importpath = "github.com/mattn/go-isatty",
        sum = "h1:BTarxUcIeDqL27Mc+vyvdWYSL28zpIhv3RoTdsLMPng=",
        version = "v0.0.17",
    )
    go_repository(
        name = "com_github_mattn_go_runewidth",
        build_file_proto_mode = "disable",
        importpath = "github.com/mattn/go-runewidth",
        sum = "h1:Lm995f3rfxdpd6TSmuVCHVb/QhupuXlYr8sCI/QdE+0=",
        version = "v0.0.9",
    )
    go_repository(
        name = "com_github_mattn_go_sqlite3",
        build_file_proto_mode = "disable",
        importpath = "github.com/mattn/go-sqlite3",
        sum = "h1:yOQRA0RpS5PFz/oikGwBEqvAWhWg5ufRz4ETLjwpU1Y=",
        version = "v1.14.16",
    )
    go_repository(
        name = "com_github_matttproud_golang_protobuf_extensions",
        build_file_proto_mode = "disable",
        importpath = "github.com/matttproud/golang_protobuf_extensions",
        sum = "h1:mmDVorXM7PCGKw94cs5zkfA9PSy5pEvNWRP0ET0TIVo=",
        version = "v1.0.4",
    )
    go_repository(
        name = "com_github_miekg_dns",
        build_file_proto_mode = "disable",
        importpath = "github.com/miekg/dns",
        sum = "h1:WMszZWJG0XmzbK9FEmzH2TVcqYzFesusSIB41b8KHxY=",
        version = "v1.1.41",
    )
    go_repository(
        name = "com_github_minio_md5_simd",
        build_file_proto_mode = "disable",
        importpath = "github.com/minio/md5-simd",
        sum = "h1:Gdi1DZK69+ZVMoNHRXJyNcxrMA4dSxoYHZSQbirFg34=",
        version = "v1.1.2",
    )
    go_repository(
        name = "com_github_minio_minio_go_v7",
        build_file_proto_mode = "disable",
        importpath = "github.com/minio/minio-go/v7",
        sum = "h1:dE5DfOtnXMXCjr/HWI6zN9vCrY6Sv666qhhiwUMvGV4=",
        version = "v7.0.49",
    )
    go_repository(
        name = "com_github_minio_sha256_simd",
        build_file_proto_mode = "disable",
        importpath = "github.com/minio/sha256-simd",
        sum = "h1:v1ta+49hkWZyvaKwrQB8elexRqm6Y0aMLjCNsrYxo6g=",
        version = "v1.0.0",
    )

    go_repository(
        name = "com_github_mitchellh_cli",
        build_file_proto_mode = "disable",
        importpath = "github.com/mitchellh/cli",
        sum = "h1:tEElEatulEHDeedTxwckzyYMA5c86fbmNIUL1hBIiTg=",
        version = "v1.1.0",
    )
    go_repository(
        name = "com_github_mitchellh_go_homedir",
        build_file_proto_mode = "disable",
        importpath = "github.com/mitchellh/go-homedir",
        sum = "h1:lukF9ziXFxDFPkA1vsr5zpc1XuPDn/wFntq5mG+4E0Y=",
        version = "v1.1.0",
    )
    go_repository(
        name = "com_github_mitchellh_go_testing_interface",
        build_file_proto_mode = "disable",
        importpath = "github.com/mitchellh/go-testing-interface",
        sum = "h1:fzU/JVNcaqHQEcVFAKeR41fkiLdIPrefOvVG1VZ96U0=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_mitchellh_go_wordwrap",
        build_file_proto_mode = "disable",
        importpath = "github.com/mitchellh/go-wordwrap",
        sum = "h1:TLuKupo69TCn6TQSyGxwI1EblZZEsQ0vMlAFQflz0v0=",
        version = "v1.0.1",
    )
    go_repository(
        name = "com_github_mitchellh_gox",
        build_file_proto_mode = "disable",
        importpath = "github.com/mitchellh/gox",
        sum = "h1:lfGJxY7ToLJQjHHwi0EX6uYBdK78egf954SQl13PQJc=",
        version = "v0.4.0",
    )
    go_repository(
        name = "com_github_mitchellh_iochan",
        build_file_proto_mode = "disable",
        importpath = "github.com/mitchellh/iochan",
        sum = "h1:C+X3KsSTLFVBr/tK1eYN/vs4rJcvsiLU338UhYPJWeY=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_mitchellh_mapstructure",
        build_file_proto_mode = "disable",
        importpath = "github.com/mitchellh/mapstructure",
        sum = "h1:jeMsZIYE/09sWLaz43PL7Gy6RuMjD2eJVyuac5Z2hdY=",
        version = "v1.5.0",
    )
    go_repository(
        name = "com_github_moby_spdystream",
        build_file_proto_mode = "disable",
        importpath = "github.com/moby/spdystream",
        sum = "h1:cjW1zVyyoiM0T7b6UoySUFqzXMoqRckQtXwGPiBhOM8=",
        version = "v0.2.0",
    )

    go_repository(
        name = "com_github_modern_go_concurrent",
        build_file_proto_mode = "disable",
        importpath = "github.com/modern-go/concurrent",
        sum = "h1:TRLaZ9cD/w8PVh93nsPXa1VrQ6jlwL5oN8l14QlcNfg=",
        version = "v0.0.0-20180306012644-bacd9c7ef1dd",
    )
    go_repository(
        name = "com_github_modern_go_reflect2",
        build_file_proto_mode = "disable",
        importpath = "github.com/modern-go/reflect2",
        sum = "h1:xBagoLtFs94CBntxluKeaWgTMpvLxC4ur3nMaC9Gz0M=",
        version = "v1.0.2",
    )
    go_repository(
        name = "com_github_mojocn_base64captcha",
        build_file_proto_mode = "disable",
        importpath = "github.com/mojocn/base64Captcha",
        sum = "h1:Qeilr7Ta6eDtG4S+tQuZ5+hO+QHbiGAJdi4PfoagaA0=",
        version = "v1.3.5",
    )
    go_repository(
        name = "com_github_montanaflynn_stats",
        build_file_proto_mode = "disable",
        importpath = "github.com/montanaflynn/stats",
        sum = "h1:Duep6KMIDpY4Yo11iFsvyqJDyfzLF9+sndUKT+v64GQ=",
        version = "v0.6.6",
    )
    go_repository(
        name = "com_github_moul_http2curl",
        build_file_proto_mode = "disable",
        importpath = "github.com/moul/http2curl",
        sum = "h1:dRMWoAtb+ePxMlLkrCbAqh4TlPHXvoGUSQ323/9Zahs=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_mozillazg_go_httpheader",
        build_file_proto_mode = "disable",
        importpath = "github.com/mozillazg/go-httpheader",
        sum = "h1:IRP+HFrMX2SlwY9riuio7raffXUpzAosHtZu25BSJok=",
        version = "v0.3.1",
    )
    go_repository(
        name = "com_github_munnerz_goautoneg",
        build_file_proto_mode = "disable",
        importpath = "github.com/munnerz/goautoneg",
        sum = "h1:C3w9PqII01/Oq1c1nUAm88MOHcQC9l5mIlSMApZMrHA=",
        version = "v0.0.0-20191010083416-a7dc8b61c822",
    )

    go_repository(
        name = "com_github_mwitkow_go_conntrack",
        build_file_proto_mode = "disable",
        importpath = "github.com/mwitkow/go-conntrack",
        sum = "h1:KUppIJq7/+SVif2QVs3tOP0zanoHgBEVAwHxUSIzRqU=",
        version = "v0.0.0-20190716064945-2f068394615f",
    )
    go_repository(
        name = "com_github_mxk_go_flowrate",
        build_file_proto_mode = "disable",
        importpath = "github.com/mxk/go-flowrate",
        sum = "h1:y5//uYreIhSUg3J1GEMiLbxo1LJaP8RfCpH6pymGZus=",
        version = "v0.0.0-20140419014527-cca7078d478f",
    )
    go_repository(
        name = "com_github_nacos_group_nacos_sdk_go",
        build_file_proto_mode = "disable",
        importpath = "github.com/nacos-group/nacos-sdk-go",
        sum = "h1:xNlSC9li2A11ifTA8HCqgM6NRImGUJA4X+gGK5muJuQ=",
        version = "v1.1.3",
    )
    go_repository(
        name = "com_github_natefinch_lumberjack",
        build_file_proto_mode = "disable",
        importpath = "github.com/natefinch/lumberjack",
        sum = "h1:4QJd3OLAMgj7ph+yZTuX13Ld4UpgHp07nNdFX7mqFfM=",
        version = "v2.0.0+incompatible",
    )

    go_repository(
        name = "com_github_nats_io_jwt",
        build_file_proto_mode = "disable",
        importpath = "github.com/nats-io/jwt",
        sum = "h1:+RB5hMpXUUA2dfxuhBTEkMOrYmM+gKIZYS1KjSostMI=",
        version = "v0.3.2",
    )
    go_repository(
        name = "com_github_nats_io_nats_go",
        build_file_proto_mode = "disable",
        importpath = "github.com/nats-io/nats.go",
        sum = "h1:ik3HbLhZ0YABLto7iX80pZLPw/6dx3T+++MZJwLnMrQ=",
        version = "v1.9.1",
    )
    go_repository(
        name = "com_github_nats_io_nats_server_v2",
        build_file_proto_mode = "disable",
        importpath = "github.com/nats-io/nats-server/v2",
        sum = "h1:i2Ly0B+1+rzNZHHWtD4ZwKi+OU5l+uQo1iDHZ2PmiIc=",
        version = "v2.1.2",
    )
    go_repository(
        name = "com_github_nats_io_nkeys",
        build_file_proto_mode = "disable",
        importpath = "github.com/nats-io/nkeys",
        sum = "h1:6JrEfig+HzTH85yxzhSVbjHRJv9cn0p6n3IngIcM5/k=",
        version = "v0.1.3",
    )
    go_repository(
        name = "com_github_nats_io_nuid",
        build_file_proto_mode = "disable",
        importpath = "github.com/nats-io/nuid",
        sum = "h1:5iA8DT8V7q8WK2EScv2padNa/rTESc1KdnPw4TC2paw=",
        version = "v1.0.1",
    )
    go_repository(
        name = "com_github_niemeyer_pretty",
        build_file_proto_mode = "disable",
        importpath = "github.com/niemeyer/pretty",
        sum = "h1:fD57ERR4JtEqsWbfPhv4DMiApHyliiK5xCTNVSPiaAs=",
        version = "v0.0.0-20200227124842-a10e7caefd8e",
    )
    go_repository(
        name = "com_github_nishanths_predeclared",
        build_file_proto_mode = "disable",
        importpath = "github.com/nishanths/predeclared",
        sum = "h1:3f0nxAmdj/VoCGN/ijdMy7bj6SBagaqYg1B0hu8clMA=",
        version = "v0.0.0-20200524104333-86fad755b4d3",
    )

    go_repository(
        name = "com_github_nxadm_tail",
        build_file_proto_mode = "disable",
        importpath = "github.com/nxadm/tail",
        sum = "h1:nPr65rt6Y5JFSKQO7qToXr7pePgD6Gwiw05lkbyAQTE=",
        version = "v1.4.8",
    )
    go_repository(
        name = "com_github_nytimes_gziphandler",
        build_file_proto_mode = "disable",
        importpath = "github.com/NYTimes/gziphandler",
        sum = "h1:lsxEuwrXEAokXB9qhlbKWPpo3KMLZQ5WB5WLQRW1uq0=",
        version = "v0.0.0-20170623195520-56545f4a5d46",
    )
    go_repository(
        name = "com_github_ogen_go_ogen",
        build_file_proto_mode = "disable",
        importpath = "github.com/ogen-go/ogen",
        sum = "h1:nvjeoT8wBu4oYVK8b4veOjBHA1puwVEE56Enx2blNDI=",
        version = "v0.56.1",
    )

    go_repository(
        name = "com_github_oklog_oklog",
        build_file_proto_mode = "disable",
        importpath = "github.com/oklog/oklog",
        sum = "h1:wVfs8F+in6nTBMkA7CbRw+zZMIB7nNM825cM1wuzoTk=",
        version = "v0.3.2",
    )
    go_repository(
        name = "com_github_oklog_run",
        build_file_proto_mode = "disable",
        importpath = "github.com/oklog/run",
        sum = "h1:Ru7dDtJNOyC66gQ5dQmaCa0qIsAUFY3sFpK1Xk8igrw=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_oklog_ulid",
        build_file_proto_mode = "disable",
        importpath = "github.com/oklog/ulid",
        sum = "h1:EGfNDEx6MqHz8B3uNV6QAib1UR2Lm97sHi3ocA6ESJ4=",
        version = "v1.3.1",
    )
    go_repository(
        name = "com_github_oklog_ulid_v2",
        build_file_proto_mode = "disable",
        importpath = "github.com/oklog/ulid/v2",
        sum = "h1:r4fFzBm+bv0wNKNh5eXTwU7i85y5x+uwkxCUTNVQqLc=",
        version = "v2.0.2",
    )

    go_repository(
        name = "com_github_olekukonko_tablewriter",
        build_file_proto_mode = "disable",
        importpath = "github.com/olekukonko/tablewriter",
        sum = "h1:P2Ga83D34wi1o9J6Wh1mRuqd4mF/x/lgBS7N7AbDhec=",
        version = "v0.0.5",
    )
    go_repository(
        name = "com_github_oneofone_xxhash",
        build_file_proto_mode = "disable",
        importpath = "github.com/OneOfOne/xxhash",
        sum = "h1:KMrpdQIwFcEqXDklaen+P1axHaj9BSKzvpUUfnHldSE=",
        version = "v1.2.2",
    )
    go_repository(
        name = "com_github_onsi_ginkgo",
        build_file_proto_mode = "disable",
        importpath = "github.com/onsi/ginkgo",
        sum = "h1:8xi0RTUf59SOSfEtZMvwTvXYMzG4gV23XVHOZiXNtnE=",
        version = "v1.16.5",
    )
    go_repository(
        name = "com_github_onsi_ginkgo_v2",
        build_file_proto_mode = "disable",
        importpath = "github.com/onsi/ginkgo/v2",
        sum = "h1:+Ig9nvqgS5OBSACXNk15PLdp0U9XPYROt9CFzVdFGIs=",
        version = "v2.4.0",
    )

    go_repository(
        name = "com_github_onsi_gomega",
        build_file_proto_mode = "disable",
        importpath = "github.com/onsi/gomega",
        sum = "h1:/oxKu9c2HVap+F3PfKort2Hw5DEU+HGlW8n+tguWsys=",
        version = "v1.23.0",
    )
    go_repository(
        name = "com_github_op_go_logging",
        build_file_proto_mode = "disable",
        importpath = "github.com/op/go-logging",
        sum = "h1:lDH9UUVJtmYCjyT0CI4q8xvlXPxeZ0gYCVvWbmPlp88=",
        version = "v0.0.0-20160315200505-970db520ece7",
    )
    go_repository(
        name = "com_github_opentracing_basictracer_go",
        build_file_proto_mode = "disable",
        importpath = "github.com/opentracing/basictracer-go",
        sum = "h1:YyUAhaEfjoWXclZVJ9sGoNct7j4TVk7lZWlQw5UXuoo=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_opentracing_contrib_go_observer",
        build_file_proto_mode = "disable",
        importpath = "github.com/opentracing-contrib/go-observer",
        sum = "h1:lM6RxxfUMrYL/f8bWEUqdXrANWtrL7Nndbm9iFN0DlU=",
        version = "v0.0.0-20170622124052-a52f23424492",
    )
    go_repository(
        name = "com_github_opentracing_opentracing_go",
        build_file_proto_mode = "disable",
        importpath = "github.com/opentracing/opentracing-go",
        sum = "h1:FfH+VrHHk6Lxt9HdVS0PXzSXFyS2NbZKXv33FYPol0A=",
        version = "v1.2.1-0.20220228012449-10b1cf09e00b",
    )
    go_repository(
        name = "com_github_openzipkin_contrib_zipkin_go_opentracing",
        build_file_proto_mode = "disable",
        importpath = "github.com/openzipkin-contrib/zipkin-go-opentracing",
        sum = "h1:ZCnq+JUrvXcDVhX/xRolRBZifmabN1HcS1wrPSvxhrU=",
        version = "v0.4.5",
    )
    go_repository(
        name = "com_github_openzipkin_zipkin_go",
        build_file_proto_mode = "disable",
        importpath = "github.com/openzipkin/zipkin-go",
        sum = "h1:kNd/ST2yLLWhaWrkgchya40TJabe8Hioj9udfPcEO5A=",
        version = "v0.4.1",
    )
    go_repository(
        name = "com_github_pact_foundation_pact_go",
        build_file_proto_mode = "disable",
        importpath = "github.com/pact-foundation/pact-go",
        sum = "h1:OYkFijGHoZAYbOIb1LWXrwKQbMMRUv1oQ89blD2Mh2Q=",
        version = "v1.0.4",
    )
    go_repository(
        name = "com_github_pascaldekloe_goe",
        build_file_proto_mode = "disable",
        importpath = "github.com/pascaldekloe/goe",
        sum = "h1:cBOtyMzM9HTpWjXfbbunk26uA6nG3a8n06Wieeh0MwY=",
        version = "v0.1.0",
    )
    go_repository(
        name = "com_github_patrickmn_go_cache",
        build_file_proto_mode = "disable",
        importpath = "github.com/patrickmn/go-cache",
        sum = "h1:HRMgzkcYKYpi3C8ajMPV8OFXaaRUnok+kx1WdO15EQc=",
        version = "v2.1.0+incompatible",
    )

    go_repository(
        name = "com_github_pborman_uuid",
        build_file_proto_mode = "disable",
        importpath = "github.com/pborman/uuid",
        sum = "h1:J7Q5mO4ysT1dv8hyrUGHb9+ooztCXu1D8MY8DZYsu3g=",
        version = "v1.2.0",
    )
    go_repository(
        name = "com_github_pelletier_go_toml",
        build_file_proto_mode = "disable",
        importpath = "github.com/pelletier/go-toml",
        sum = "h1:4yBQzkHv+7BHq2PQUZF3Mx0IYxG7LsP222s7Agd3ve8=",
        version = "v1.9.5",
    )
    go_repository(
        name = "com_github_pelletier_go_toml_v2",
        build_file_proto_mode = "disable",
        importpath = "github.com/pelletier/go-toml/v2",
        sum = "h1:nrzqCb7j9cDFj2coyLNLaZuJTLjWjlaz6nvTvIwycIU=",
        version = "v2.0.6",
    )
    go_repository(
        name = "com_github_performancecopilot_speed",
        build_file_proto_mode = "disable",
        importpath = "github.com/performancecopilot/speed",
        sum = "h1:2WnRzIquHa5QxaJKShDkLM+sc0JPuwhXzK8OYOyt3Vg=",
        version = "v3.0.0+incompatible",
    )
    go_repository(
        name = "com_github_peterbourgon_diskv",
        build_file_proto_mode = "disable",
        importpath = "github.com/peterbourgon/diskv",
        sum = "h1:UBdAOUP5p4RWqPBg048CAvpKN+vxiaj6gdUUzhl4XmI=",
        version = "v2.0.1+incompatible",
    )
    go_repository(
        name = "com_github_philhofer_fwd",
        build_file_proto_mode = "disable",
        importpath = "github.com/philhofer/fwd",
        sum = "h1:GdGcTjf5RNAxwS4QLsiMzJYj5KEvPJD3Abr261yRQXQ=",
        version = "v1.1.1",
    )

    go_repository(
        name = "com_github_pierrec_lz4",
        build_file_proto_mode = "disable",
        importpath = "github.com/pierrec/lz4",
        sum = "h1:9UY3+iC23yxF0UfGaYrGplQ+79Rg+h/q9FV9ix19jjM=",
        version = "v2.6.1+incompatible",
    )
    go_repository(
        name = "com_github_pierrec_lz4_v4",
        build_file_proto_mode = "disable",
        importpath = "github.com/pierrec/lz4/v4",
        sum = "h1:kV4Ip+/hUBC+8T6+2EgburRtkE9ef4nbY3f4dFhGjMc=",
        version = "v4.1.17",
    )
    go_repository(
        name = "com_github_pkg_diff",
        build_file_proto_mode = "disable",
        importpath = "github.com/pkg/diff",
        sum = "h1:aoZm08cpOy4WuID//EZDgcC4zIxODThtZNPirFr42+A=",
        version = "v0.0.0-20210226163009-20ebb0f2a09e",
    )
    go_repository(
        name = "com_github_pkg_errors",
        build_file_proto_mode = "disable",
        importpath = "github.com/pkg/errors",
        sum = "h1:FEBLx1zS214owpjy7qsBeixbURkuhQAwrK5UwLGTwt4=",
        version = "v0.9.1",
    )
    go_repository(
        name = "com_github_pkg_profile",
        build_file_proto_mode = "disable",
        importpath = "github.com/pkg/profile",
        sum = "h1:F++O52m40owAmADcojzM+9gyjmMOY/T4oYJkgFDH8RE=",
        version = "v1.2.1",
    )
    go_repository(
        name = "com_github_pkg_sftp",
        build_file_proto_mode = "disable",
        importpath = "github.com/pkg/sftp",
        sum = "h1:I2qBYMChEhIjOgazfJmV3/mZM256btk6wkCDRmW7JYs=",
        version = "v1.13.1",
    )

    go_repository(
        name = "com_github_pmezard_go_difflib",
        build_file_proto_mode = "disable",
        importpath = "github.com/pmezard/go-difflib",
        sum = "h1:4DBwDE0NGyQoBHbLQYPwSUPoCMWR5BEzIk/f1lZbAQM=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_polarismesh_polaris_go",
        build_file_proto_mode = "disable",
        importpath = "github.com/polarismesh/polaris-go",
        sum = "h1:KZKX//ow4OPPoS5+s7h07ptprg+2AcNVGrN6WakC9QM=",
        version = "v1.3.0",
    )

    go_repository(
        name = "com_github_posener_complete",
        build_file_proto_mode = "disable",
        importpath = "github.com/posener/complete",
        sum = "h1:NP0eAhjcjImqslEwo/1hq7gpajME0fTLTezBKDqfXqo=",
        version = "v1.2.3",
    )
    go_repository(
        name = "com_github_power_devops_perfstat",
        build_file_proto_mode = "disable",
        importpath = "github.com/power-devops/perfstat",
        sum = "h1:0LFwY6Q3gMACTjAbMZBjXAqTOzOwFaj2Ld6cjeQ7Rig=",
        version = "v0.0.0-20221212215047-62379fc7944b",
    )
    go_repository(
        name = "com_github_prashantv_gostub",
        build_file_proto_mode = "disable",
        importpath = "github.com/prashantv/gostub",
        sum = "h1:BTyx3RfQjRHnUWaGF9oQos79AlQ5k8WNktv7VGvVH4g=",
        version = "v1.1.0",
    )

    go_repository(
        name = "com_github_prometheus_client_golang",
        build_file_proto_mode = "disable",
        importpath = "github.com/prometheus/client_golang",
        sum = "h1:nJdhIvne2eSX/XRAFV9PcvFFRbrjbcTUj0VP62TMhnw=",
        version = "v1.14.0",
    )
    go_repository(
        name = "com_github_prometheus_client_model",
        build_file_proto_mode = "disable",
        importpath = "github.com/prometheus/client_model",
        sum = "h1:UBgGFHqYdG/TPFD1B1ogZywDqEkwp3fBMvqdiQ7Xew4=",
        version = "v0.3.0",
    )
    go_repository(
        name = "com_github_prometheus_common",
        build_file_proto_mode = "disable",
        importpath = "github.com/prometheus/common",
        sum = "h1:oOyhkDq05hPZKItWVBkJ6g6AtGxi+fy7F4JvUV8uhsI=",
        version = "v0.39.0",
    )
    go_repository(
        name = "com_github_prometheus_procfs",
        build_file_proto_mode = "disable",
        importpath = "github.com/prometheus/procfs",
        sum = "h1:wzCHvIvM5SxWqYvwgVL7yJY8Lz3PKn49KQtpgMYJfhI=",
        version = "v0.9.0",
    )
    go_repository(
        name = "com_github_prometheus_tsdb",
        build_file_proto_mode = "disable",
        importpath = "github.com/prometheus/tsdb",
        sum = "h1:YZcsG11NqnK4czYLrWd9mpEuAJIHVQLwdrleYfszMAA=",
        version = "v0.7.1",
    )
    go_repository(
        name = "com_github_puerkitobio_purell",
        build_file_proto_mode = "disable",
        importpath = "github.com/PuerkitoBio/purell",
        sum = "h1:WEQqlqaGbrPkxLJWfBwQmfEAE1Z7ONdDLqrN38tNFfI=",
        version = "v1.1.1",
    )
    go_repository(
        name = "com_github_puerkitobio_urlesc",
        build_file_proto_mode = "disable",
        importpath = "github.com/PuerkitoBio/urlesc",
        sum = "h1:d+Bc7a5rLufV/sSk/8dngufqelfh6jnri85riMAaF/M=",
        version = "v0.0.0-20170810143723-de5bf2ad4578",
    )

    go_repository(
        name = "com_github_qcloudapi_qcloud_sign_golang",
        build_file_proto_mode = "disable",
        importpath = "github.com/QcloudApi/qcloud_sign_golang",
        sum = "h1:DTQ/38ao/CfXsrK0cSAL+h4R/u0VVvfWLZEOlLwEROI=",
        version = "v0.0.0-20141224014652-e4130a326409",
    )
    go_repository(
        name = "com_github_qiniu_dyn",
        build_file_proto_mode = "disable",
        importpath = "github.com/qiniu/dyn",
        sum = "h1:s+xPTeV0H8yikgM4ZMBc7Rrefam8UNI3asBlkaOQg5o=",
        version = "v1.3.0",
    )
    go_repository(
        name = "com_github_qiniu_go_sdk_v7",
        build_file_proto_mode = "disable",
        importpath = "github.com/qiniu/go-sdk/v7",
        sum = "h1:6icihMTKHoKMmeU1mqtIoHUv7c1LrLjYm8wTQaYDqmw=",
        version = "v7.14.0",
    )
    go_repository(
        name = "com_github_qiniu_x",
        build_file_proto_mode = "disable",
        importpath = "github.com/qiniu/x",
        sum = "h1:7V/CYWEmo9axJULvrJN6sMYh2FdY+esN5h8jwDkA4b0=",
        version = "v1.10.5",
    )
    go_repository(
        name = "com_github_rabbitmq_amqp091_go",
        build_file_proto_mode = "disable",
        importpath = "github.com/rabbitmq/amqp091-go",
        sum = "h1:VouyHPBu1CrKyJVfteGknGOGCzmOz0zcv/tONLkb7rg=",
        version = "v1.5.0",
    )

    go_repository(
        name = "com_github_rakyll_statik",
        build_file_proto_mode = "disable",
        importpath = "github.com/rakyll/statik",
        sum = "h1:OF3QCZUuyPxuGEP7B4ypUa7sB/iHtqOTDYZXGM8KOdQ=",
        version = "v0.1.7",
    )
    go_repository(
        name = "com_github_rcrowley_go_metrics",
        build_file_proto_mode = "disable",
        importpath = "github.com/rcrowley/go-metrics",
        sum = "h1:N/ElC8H3+5XpJzTSTfLsJV/mx9Q9g7kxmchpfZyxgzM=",
        version = "v0.0.0-20201227073835-cf1acfcdf475",
    )
    go_repository(
        name = "com_github_rogpeppe_fastuuid",
        build_file_proto_mode = "disable",
        importpath = "github.com/rogpeppe/fastuuid",
        sum = "h1:Ppwyp6VYCF1nvBTXL3trRso7mXMlRrw9ooo375wvi2s=",
        version = "v1.2.0",
    )
    go_repository(
        name = "com_github_rogpeppe_go_internal",
        build_file_proto_mode = "disable",
        importpath = "github.com/rogpeppe/go-internal",
        sum = "h1:73kH8U+JUqXU8lRuOHeVHaa/SZPifC7BkcraZVejAe8=",
        version = "v1.9.0",
    )
    go_repository(
        name = "com_github_rs_xid",
        build_file_proto_mode = "disable",
        importpath = "github.com/rs/xid",
        sum = "h1:qd7wPTDkN6KQx2VmMBLrpHkiyQwgFXRnkOLacUiaSNY=",
        version = "v1.4.0",
    )
    go_repository(
        name = "com_github_rs_zerolog",
        build_file_proto_mode = "disable",
        importpath = "github.com/rs/zerolog",
        sum = "h1:uPRuwkWF4J6fGsJ2R0Gn2jB1EQiav9k3S6CSdygQJXY=",
        version = "v1.15.0",
    )
    go_repository(
        name = "com_github_russross_blackfriday_v2",
        build_file_proto_mode = "disable",
        importpath = "github.com/russross/blackfriday/v2",
        sum = "h1:lPqVAte+HuHNfhJ/0LC98ESWRz8afy9tM/0RK8m9o+Q=",
        version = "v2.0.1",
    )
    go_repository(
        name = "com_github_ryanuber_columnize",
        build_file_proto_mode = "disable",
        importpath = "github.com/ryanuber/columnize",
        sum = "h1:j1Wcmh8OrK4Q7GXY+V7SVSY8nUWQxHW5TkBe7YUl+2s=",
        version = "v2.1.0+incompatible",
    )
    go_repository(
        name = "com_github_sagikazarmark_crypt",
        build_file_proto_mode = "disable",
        importpath = "github.com/sagikazarmark/crypt",
        sum = "h1:xtk0uUHVWVsRBdEUGYBym4CXbcllXky2M7Qlwsf8C0Y=",
        version = "v0.8.0",
    )

    go_repository(
        name = "com_github_samuel_go_zookeeper",
        build_file_proto_mode = "disable",
        importpath = "github.com/samuel/go-zookeeper",
        sum = "h1:p3Vo3i64TCLY7gIfzeQaUJ+kppEO5WQG3cL8iE8tGHU=",
        version = "v0.0.0-20190923202752-2cc03de413da",
    )
    go_repository(
        name = "com_github_satori_go_uuid",
        build_file_proto_mode = "disable",
        importpath = "github.com/satori/go.uuid",
        sum = "h1:0uYX9dsZ2yD7q2RtLRtPSdGDWzjeM3TbMJP9utgA0ww=",
        version = "v1.2.0",
    )
    go_repository(
        name = "com_github_sclevine_agouti",
        build_file_proto_mode = "disable",
        importpath = "github.com/sclevine/agouti",
        sum = "h1:8IBJS6PWz3uTlMP3YBIR5f+KAldcGuOeFkFbUWfBgK4=",
        version = "v3.0.0+incompatible",
    )
    go_repository(
        name = "com_github_sean_seed",
        build_file_proto_mode = "disable",
        importpath = "github.com/sean-/seed",
        sum = "h1:nn5Wsu0esKSJiIVhscUtVbo7ada43DJhG55ua/hjS5I=",
        version = "v0.0.0-20170313163322-e2103e2c3529",
    )
    go_repository(
        name = "com_github_segmentio_asm",
        build_file_proto_mode = "disable",
        importpath = "github.com/segmentio/asm",
        sum = "h1:9BQrFxC+YOHJlTlHGkTrFWf59nbL3XnCoFLTwDCI7ys=",
        version = "v1.2.0",
    )

    go_repository(
        name = "com_github_segmentio_kafka_go",
        build_file_proto_mode = "disable",
        importpath = "github.com/segmentio/kafka-go",
        sum = "h1:iQdOBbUSdfuYlFpvjuALgj7N6DrdPA0HfB4AhREOdtg=",
        version = "v0.4.38",
    )

    go_repository(
        name = "com_github_sergi_go_diff",
        build_file_proto_mode = "disable",
        importpath = "github.com/sergi/go-diff",
        sum = "h1:we8PVUC3FE2uYfodKH/nBHMSetSfHDR6scGdBi+erh0=",
        version = "v1.1.0",
    )
    go_repository(
        name = "com_github_shirou_gopsutil_v3",
        build_file_proto_mode = "disable",
        importpath = "github.com/shirou/gopsutil/v3",
        sum = "h1:kxsPKS+Eeo+VnEQ2XCaGJepeP6KY53QoRTETx3+1ndM=",
        version = "v3.22.11",
    )
    go_repository(
        name = "com_github_shonminh_apollo_client",
        build_file_proto_mode = "disable",
        importpath = "github.com/Shonminh/apollo-client",
        sum = "h1:AXGp4wOahrEKjheMXehgsG9B8dEfLQHttRLHeEefvus=",
        version = "v0.4.0",
    )

    go_repository(
        name = "com_github_shopify_sarama",
        build_file_proto_mode = "disable",
        importpath = "github.com/Shopify/sarama",
        sum = "h1:lqqPUPQZ7zPqYlWpTh+LQ9bhYNu2xJL6k1SJN4WVe2A=",
        version = "v1.38.1",
    )
    go_repository(
        name = "com_github_shopify_toxiproxy",
        build_file_proto_mode = "disable",
        importpath = "github.com/Shopify/toxiproxy",
        sum = "h1:TKdv8HiTLgE5wdJuEML90aBgNWsokNbMijUGhmcoBJc=",
        version = "v2.1.4+incompatible",
    )
    go_repository(
        name = "com_github_shopify_toxiproxy_v2",
        build_file_proto_mode = "disable",
        importpath = "github.com/Shopify/toxiproxy/v2",
        sum = "h1:i4LPT+qrSlKNtQf5QliVjdP08GyAH8+BUIc9gT0eahc=",
        version = "v2.5.0",
    )
    go_repository(
        name = "com_github_shopspring_decimal",
        build_file_proto_mode = "disable",
        importpath = "github.com/shopspring/decimal",
        sum = "h1:abSATXmQEYyShuxI4/vyW3tV1MrKAJzCZ/0zLUXYbsQ=",
        version = "v1.2.0",
    )
    go_repository(
        name = "com_github_shurcool_sanitized_anchor_name",
        build_file_proto_mode = "disable",
        importpath = "github.com/shurcooL/sanitized_anchor_name",
        sum = "h1:PdmoCO6wvbs+7yrJyMORt4/BmY5IYyJwS/kOiWx8mHo=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_sirupsen_logrus",
        build_file_proto_mode = "disable",
        importpath = "github.com/sirupsen/logrus",
        sum = "h1:trlNQbNUG3OdDrDil03MCb1H2o9nJ1x4/5LYw7byDE0=",
        version = "v1.9.0",
    )
    go_repository(
        name = "com_github_smartystreets_assertions",
        build_file_proto_mode = "disable",
        importpath = "github.com/smartystreets/assertions",
        sum = "h1:zE9ykElWQ6/NYmHa3jpm/yHnI4xSofP+UP6SpjHcSeM=",
        version = "v0.0.0-20180927180507-b2de0cb4f26d",
    )
    go_repository(
        name = "com_github_smartystreets_goconvey",
        build_file_proto_mode = "disable",
        importpath = "github.com/smartystreets/goconvey",
        sum = "h1:fv0U8FUIMPNf1L9lnHLvLhgicrIVChEkdzIKYqbNC9s=",
        version = "v1.6.4",
    )
    go_repository(
        name = "com_github_soheilhy_cmux",
        build_file_proto_mode = "disable",
        importpath = "github.com/soheilhy/cmux",
        sum = "h1:jjzc5WVemNEDTLwv9tlmemhC73tI08BNOIGwBOo10Js=",
        version = "v0.1.5",
    )
    go_repository(
        name = "com_github_sony_gobreaker",
        build_file_proto_mode = "disable",
        importpath = "github.com/sony/gobreaker",
        sum = "h1:oMnRNZXX5j85zso6xCPRNPtmAycat+WcoKbklScLDgQ=",
        version = "v0.4.1",
    )
    go_repository(
        name = "com_github_sony_sonyflake",
        build_file_proto_mode = "disable",
        importpath = "github.com/sony/sonyflake",
        sum = "h1:wnrEcL3aOkWmPlhScLEGAXKkLAIslnBteNUq4Bw6MM4=",
        version = "v1.1.0",
    )
    go_repository(
        name = "com_github_spaolacci_murmur3",
        build_file_proto_mode = "disable",
        importpath = "github.com/spaolacci/murmur3",
        sum = "h1:7c1g84S4BPRrfL5Xrdp6fOJ206sU9y293DDHaoy0bLI=",
        version = "v1.1.0",
    )
    go_repository(
        name = "com_github_spf13_afero",
        build_file_proto_mode = "disable",
        importpath = "github.com/spf13/afero",
        sum = "h1:j49Hj62F0n+DaZ1dDCvhABaPNSGNkt32oRFxI33IEMw=",
        version = "v1.9.2",
    )
    go_repository(
        name = "com_github_spf13_cast",
        build_file_proto_mode = "disable",
        importpath = "github.com/spf13/cast",
        sum = "h1:rj3WzYc11XZaIZMPKmwP96zkFEnnAmV8s6XbB2aY32w=",
        version = "v1.5.0",
    )
    go_repository(
        name = "com_github_spf13_cobra",
        build_file_proto_mode = "disable",
        importpath = "github.com/spf13/cobra",
        sum = "h1:o94oiPyS4KD1mPy2fmcYYHHfCxLqYjJOhGsCHFZtEzA=",
        version = "v1.6.1",
    )
    go_repository(
        name = "com_github_spf13_jwalterweatherman",
        build_file_proto_mode = "disable",
        importpath = "github.com/spf13/jwalterweatherman",
        sum = "h1:ue6voC5bR5F8YxI5S67j9i582FU4Qvo2bmqnqMYADFk=",
        version = "v1.1.0",
    )
    go_repository(
        name = "com_github_spf13_pflag",
        build_file_proto_mode = "disable",
        importpath = "github.com/spf13/pflag",
        sum = "h1:iy+VFUOCP1a+8yFto/drg2CJ5u0yRoB7fZw3DKv/JXA=",
        version = "v1.0.5",
    )
    go_repository(
        name = "com_github_spf13_viper",
        build_file_proto_mode = "disable",
        importpath = "github.com/spf13/viper",
        sum = "h1:Rg7d3Lo706X9tHsJMUjdiwMpHB7W8WnSVOssIY+JElU=",
        version = "v1.14.0",
    )
    go_repository(
        name = "com_github_stackexchange_wmi",
        build_file_proto_mode = "disable",
        importpath = "github.com/StackExchange/wmi",
        sum = "h1:VIkavFPXSjcnS+O8yTq7NI32k0R5Aj+v39y29VYDOSA=",
        version = "v1.2.1",
    )
    go_repository(
        name = "com_github_stoewer_go_strcase",
        build_file_proto_mode = "disable",
        importpath = "github.com/stoewer/go-strcase",
        sum = "h1:Z2iHWqGXH00XYgqDmNgQbIBxf3wrNq0F3feEy0ainaU=",
        version = "v1.2.0",
    )

    go_repository(
        name = "com_github_streadway_amqp",
        build_file_proto_mode = "disable",
        importpath = "github.com/streadway/amqp",
        sum = "h1:WhxRHzgeVGETMlmVfqhRn8RIeeNoPr2Czh33I4Zdccw=",
        version = "v0.0.0-20190827072141-edfb9018d271",
    )
    go_repository(
        name = "com_github_streadway_handy",
        build_file_proto_mode = "disable",
        importpath = "github.com/streadway/handy",
        sum = "h1:AhmOdSHeswKHBjhsLs/7+1voOxT+LLrSk/Nxvk35fug=",
        version = "v0.0.0-20190108123426-d5acb3125c2a",
    )
    go_repository(
        name = "com_github_stretchr_objx",
        build_file_proto_mode = "disable",
        importpath = "github.com/stretchr/objx",
        sum = "h1:1zr/of2m5FGMsad5YfcqgdqdWrIhu+EBEJRhR1U7z/c=",
        version = "v0.5.0",
    )
    go_repository(
        name = "com_github_stretchr_testify",
        build_file_proto_mode = "disable",
        importpath = "github.com/stretchr/testify",
        sum = "h1:w7B6lhMri9wdJUVmEZPGGhZzrYTPvgJArz7wNPgYKsk=",
        version = "v1.8.1",
    )
    go_repository(
        name = "com_github_subosito_gotenv",
        build_file_proto_mode = "disable",
        importpath = "github.com/subosito/gotenv",
        sum = "h1:jyEFiXpy21Wm81FBN71l9VoMMV8H8jG+qIK3GCpY6Qs=",
        version = "v1.4.1",
    )
    go_repository(
        name = "com_github_tebeka_strftime",
        build_file_proto_mode = "disable",
        importpath = "github.com/tebeka/strftime",
        sum = "h1:5HQXOqWKYRFfNyBMNVc9z5+QzuBtIXy03psIhtdJYto=",
        version = "v0.1.3",
    )
    go_repository(
        name = "com_github_tencentcloud_tencentcloud_cls_sdk_go",
        build_file_proto_mode = "disable",
        importpath = "github.com/tencentcloud/tencentcloud-cls-sdk-go",
        sum = "h1:29QoDxUqlFWd0Pgyt4lqGTGMNvJGxVQ3lRx1haRaPnw=",
        version = "v1.0.2",
    )

    go_repository(
        name = "com_github_tencentcloud_tencentcloud_sdk_go_tencentcloud_common",
        build_file_proto_mode = "disable",
        importpath = "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common",
        sum = "h1:Oho9ykiKXwOHkeq5jSAvlkBAcRwNqnrUca/5WacvH2E=",
        version = "v1.0.194",
    )
    go_repository(
        name = "com_github_tencentcloud_tencentcloud_sdk_go_tencentcloud_kms",
        build_file_proto_mode = "disable",
        importpath = "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/kms",
        sum = "h1:YB6qJyCPuwtHFr54/GAfYj1VfwhiDHnwtOKu40OaG2M=",
        version = "v1.0.194",
    )
    go_repository(
        name = "com_github_tencentyun_cos_go_sdk_v5",
        build_file_proto_mode = "disable",
        importpath = "github.com/tencentyun/cos-go-sdk-v5",
        sum = "h1:W6vDGKCHe4wBACI1d2UgE6+50sJFhRWU4O8IB2ozzxM=",
        version = "v0.7.40",
    )
    go_repository(
        name = "com_github_tevid_gohamcrest",
        build_file_proto_mode = "disable",
        importpath = "github.com/tevid/gohamcrest",
        sum = "h1:ou+xSqlIw1xfGTg1uq1nif/htZ2S3EzRqLm2BP+tYU0=",
        version = "v1.1.1",
    )

    go_repository(
        name = "com_github_tidwall_assert",
        build_file_proto_mode = "disable",
        importpath = "github.com/tidwall/assert",
        sum = "h1:aWcKyRBUAdLoVebxo95N7+YZVTFF/ASTr7BN4sLP6XI=",
        version = "v0.1.0",
    )
    go_repository(
        name = "com_github_tidwall_btree",
        build_file_proto_mode = "disable",
        importpath = "github.com/tidwall/btree",
        sum = "h1:LDZfKfQIBHGHWSwckhXI0RPSXzlo+KYdjK7FWSqOzzg=",
        version = "v1.6.0",
    )
    go_repository(
        name = "com_github_tidwall_buntdb",
        build_file_proto_mode = "disable",
        importpath = "github.com/tidwall/buntdb",
        sum = "h1:U/ebfkmYPBnyiNZIirUiWFcxA/mgzjbKlyPynFsPtyM=",
        version = "v1.2.10",
    )
    go_repository(
        name = "com_github_tidwall_gjson",
        build_file_proto_mode = "disable",
        importpath = "github.com/tidwall/gjson",
        sum = "h1:uo0p8EbA09J7RQaflQ1aBRffTR7xedD2bcIVSYxLnkM=",
        version = "v1.14.4",
    )
    go_repository(
        name = "com_github_tidwall_grect",
        build_file_proto_mode = "disable",
        importpath = "github.com/tidwall/grect",
        sum = "h1:dA3oIgNgWdSspFzn1kS4S/RDpZFLrIxAZOdJKjYapOg=",
        version = "v0.1.4",
    )
    go_repository(
        name = "com_github_tidwall_lotsa",
        build_file_proto_mode = "disable",
        importpath = "github.com/tidwall/lotsa",
        sum = "h1:dNVBH5MErdaQ/xd9s769R31/n2dXavsQ0Yf4TMEHHw8=",
        version = "v1.0.2",
    )
    go_repository(
        name = "com_github_tidwall_match",
        build_file_proto_mode = "disable",
        importpath = "github.com/tidwall/match",
        sum = "h1:+Ho715JplO36QYgwN9PGYNhgZvoUSc9X2c80KVTi+GA=",
        version = "v1.1.1",
    )
    go_repository(
        name = "com_github_tidwall_pretty",
        build_file_proto_mode = "disable",
        importpath = "github.com/tidwall/pretty",
        sum = "h1:qjsOFOWWQl+N3RsoF5/ssm1pHmJJwhjlSbZ51I6wMl4=",
        version = "v1.2.1",
    )
    go_repository(
        name = "com_github_tidwall_rtred",
        build_file_proto_mode = "disable",
        importpath = "github.com/tidwall/rtred",
        sum = "h1:exmoQtOLvDoO8ud++6LwVsAMTu0KPzLTUrMln8u1yu8=",
        version = "v0.1.2",
    )
    go_repository(
        name = "com_github_tidwall_rtree",
        build_file_proto_mode = "disable",
        importpath = "github.com/tidwall/rtree",
        sum = "h1:+NL1GDIUOKxVfbp2KoJQD9cTQ6dyP2co9q4yzmT9FZo=",
        version = "v0.0.0-20180113144539-6cd427091e0e",
    )
    go_repository(
        name = "com_github_tidwall_tinyqueue",
        build_file_proto_mode = "disable",
        importpath = "github.com/tidwall/tinyqueue",
        sum = "h1:SpNEvEggbpyN5DIReaJ2/1ndroY8iyEGxPYxoSaymYE=",
        version = "v0.1.1",
    )
    go_repository(
        name = "com_github_tinylib_msgp",
        build_file_proto_mode = "disable",
        importpath = "github.com/tinylib/msgp",
        sum = "h1:i+SbKraHhnrf9M5MYmvQhFnbLhAXSDWF8WWsuyRdocw=",
        version = "v1.1.6",
    )

    go_repository(
        name = "com_github_tklauser_go_sysconf",
        build_file_proto_mode = "disable",
        importpath = "github.com/tklauser/go-sysconf",
        sum = "h1:89WgdJhk5SNwJfu+GKyYveZ4IaJ7xAkecBo+KdJV0CM=",
        version = "v0.3.11",
    )
    go_repository(
        name = "com_github_tklauser_numcpus",
        build_file_proto_mode = "disable",
        importpath = "github.com/tklauser/numcpus",
        sum = "h1:kebhY2Qt+3U6RNK7UqpYNA+tJ23IBEGKkB7JQBfDYms=",
        version = "v0.6.0",
    )
    go_repository(
        name = "com_github_tmc_grpc_websocket_proxy",
        build_file_proto_mode = "disable",
        importpath = "github.com/tmc/grpc-websocket-proxy",
        sum = "h1:uruHq4dN7GR16kFc5fp3d1RIYzJW5onx8Ybykw2YQFA=",
        version = "v0.0.0-20201229170055-e5319fda7802",
    )
    go_repository(
        name = "com_github_toolkits_concurrent",
        build_file_proto_mode = "disable",
        importpath = "github.com/toolkits/concurrent",
        sum = "h1:kF/7m/ZU+0D4Jj5eZ41Zm3IH/J8OElK1Qtd7tVKAwLk=",
        version = "v0.0.0-20150624120057-a4371d70e3e3",
    )

    go_repository(
        name = "com_github_tv42_httpunix",
        build_file_proto_mode = "disable",
        importpath = "github.com/tv42/httpunix",
        sum = "h1:G3dpKMzFDjgEh2q1Z7zUUtKa8ViPtH+ocF0bE0g00O8=",
        version = "v0.0.0-20150427012821-b75d8614f926",
    )
    go_repository(
        name = "com_github_tx7do_kratos_transport",
        build_file_proto_mode = "disable",
        importpath = "github.com/tx7do/kratos-transport",
        sum = "h1:R7WU25QScsHMcDHaXmo/P3RaVuzaE1tOHqlhc/ekvvI=",
        version = "v1.0.1",
    )
    go_repository(
        name = "com_github_tx7do_kratos_transport_broker_kafka",
        build_file_proto_mode = "disable",
        importpath = "github.com/tx7do/kratos-transport/broker/kafka",
        sum = "h1:OTsEhxNXohTTJCpt66fqTvdsG1L3Lq7UTUA6nACOJg0=",
        version = "v1.0.2",
    )
    go_repository(
        name = "com_github_tx7do_kratos_transport_transport_kafka",
        build_file_proto_mode = "disable",
        importpath = "github.com/tx7do/kratos-transport/transport/kafka",
        sum = "h1:m+tMqerXqQSxvUq4UcqoRKPF0J1Z5/b0sqDUfOWqa3Q=",
        version = "v1.0.2",
    )

    go_repository(
        name = "com_github_uber_jaeger_client_go",
        build_file_proto_mode = "disable",
        importpath = "github.com/uber/jaeger-client-go",
        sum = "h1:D6wyKGCecFaSRUpo8lCVbaOOb6ThwMmTEbhRwtKR97o=",
        version = "v2.30.0+incompatible",
    )
    go_repository(
        name = "com_github_uber_jaeger_lib",
        build_file_proto_mode = "disable",
        importpath = "github.com/uber/jaeger-lib",
        sum = "h1:td4jdvLcExb4cBISKIpHuGoVXh+dVKhn2Um6rjCsSsg=",
        version = "v2.4.1+incompatible",
    )
    go_repository(
        name = "com_github_ugorji_go",
        build_file_proto_mode = "disable",
        importpath = "github.com/ugorji/go",
        sum = "h1:j4s+tAvLfL3bZyefP2SEWmhBzmuIlH/eqNuPdFPgngw=",
        version = "v1.1.4",
    )

    go_repository(
        name = "com_github_ugorji_go_codec",
        build_file_proto_mode = "disable",
        importpath = "github.com/ugorji/go/codec",
        sum = "h1:sgBJS6COt0b/P40VouWKdseidkDgHxYGm0SAglUHfP0=",
        version = "v1.2.8",
    )
    go_repository(
        name = "com_github_urfave_cli",
        build_file_proto_mode = "disable",
        importpath = "github.com/urfave/cli",
        sum = "h1:+mkCCcOFKPnCmVYVcURKps1Xe+3zP90gSYGNfRkjoIY=",
        version = "v1.22.1",
    )
    go_repository(
        name = "com_github_valyala_bytebufferpool",
        build_file_proto_mode = "disable",
        importpath = "github.com/valyala/bytebufferpool",
        sum = "h1:GqA5TC/0021Y/b9FG4Oi9Mr3q7XYx6KllzawFIhcdPw=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_valyala_fasthttp",
        build_file_proto_mode = "disable",
        importpath = "github.com/valyala/fasthttp",
        sum = "h1:d3AAQJ2DRcxJYHm7OXNXtXt2as1vMDfxeIcFvhmGGm4=",
        version = "v1.34.0",
    )
    go_repository(
        name = "com_github_valyala_tcplisten",
        build_file_proto_mode = "disable",
        importpath = "github.com/valyala/tcplisten",
        sum = "h1:rBHj/Xf+E1tRGZyWIWwJDiRY0zc1Js+CV5DqwacVSA8=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_vektah_gqlparser_v2",
        build_file_proto_mode = "disable",
        importpath = "github.com/vektah/gqlparser/v2",
        sum = "h1:96uWUPr8zxRRakTqu0yMKCnTPEsZ5oViVDL1XvkP4Tc=",
        version = "v2.4.3-0.20220508162109-d3d9eb001575",
    )

    go_repository(
        name = "com_github_vividcortex_gohistogram",
        build_file_proto_mode = "disable",
        importpath = "github.com/VividCortex/gohistogram",
        sum = "h1:6+hBz+qvs0JOrrNhhmR7lFxo5sINxBCGXrdtl/UvroE=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_vmihailenco_msgpack_v4",
        build_file_proto_mode = "disable",
        importpath = "github.com/vmihailenco/msgpack/v4",
        sum = "h1:07s4sz9IReOgdikxLTKNbBdqDMLsjPKXwvCazn8G65U=",
        version = "v4.3.12",
    )
    go_repository(
        name = "com_github_vmihailenco_msgpack_v5",
        build_file_proto_mode = "disable",
        importpath = "github.com/vmihailenco/msgpack/v5",
        sum = "h1:iBRIniTnWOo0kqkg3k3XR8Vn6OCkVlIuZNo0UoBrKx4=",
        version = "v5.0.0-beta.9",
    )

    go_repository(
        name = "com_github_vmihailenco_tagparser",
        build_file_proto_mode = "disable",
        importpath = "github.com/vmihailenco/tagparser",
        sum = "h1:gnjoVuB/kljJ5wICEEOpx98oXMWPLj22G67Vbd1qPqc=",
        version = "v0.1.2",
    )
    go_repository(
        name = "com_github_wsxiaoys_terminal",
        build_file_proto_mode = "disable",
        importpath = "github.com/wsxiaoys/terminal",
        sum = "h1:3UeQBvD0TFrlVjOeLOBz+CPAI8dnbqNSVwUwRrkp7vQ=",
        version = "v0.0.0-20160513160801-0940f3fc43a0",
    )

    go_repository(
        name = "com_github_xdg_go_pbkdf2",
        build_file_proto_mode = "disable",
        importpath = "github.com/xdg-go/pbkdf2",
        sum = "h1:Su7DPu48wXMwC3bs7MCNG+z4FhcyEuz5dlvchbq0B0c=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_xdg_go_scram",
        build_file_proto_mode = "disable",
        importpath = "github.com/xdg-go/scram",
        sum = "h1:FHX5I5B4i4hKRVRBCFRxq1iQRej7WO3hhBuJf+UUySY=",
        version = "v1.1.2",
    )
    go_repository(
        name = "com_github_xdg_go_stringprep",
        build_file_proto_mode = "disable",
        importpath = "github.com/xdg-go/stringprep",
        sum = "h1:XLI/Ng3O1Atzq0oBs3TWm+5ZVgkq2aqdlvP9JtoZ6c8=",
        version = "v1.0.4",
    )
    go_repository(
        name = "com_github_xdg_scram",
        build_file_proto_mode = "disable",
        importpath = "github.com/xdg/scram",
        sum = "h1:TuS0RFmt5Is5qm9Tm2SoD89OPqe4IRiFtyFY4iwWXsw=",
        version = "v1.0.5",
    )
    go_repository(
        name = "com_github_xdg_stringprep",
        build_file_proto_mode = "disable",
        importpath = "github.com/xdg/stringprep",
        sum = "h1:cmL5Enob4W83ti/ZHuZLuKD/xqJfus4fVPwE+/BDm+4=",
        version = "v1.0.3",
    )

    go_repository(
        name = "com_github_xeipuuv_gojsonpointer",
        build_file_proto_mode = "disable",
        importpath = "github.com/xeipuuv/gojsonpointer",
        sum = "h1:J9EGpcZtP0E/raorCMxlFGSTBrsSlaDGf3jU/qvAE2c=",
        version = "v0.0.0-20180127040702-4e3ac2762d5f",
    )
    go_repository(
        name = "com_github_xeipuuv_gojsonreference",
        build_file_proto_mode = "disable",
        importpath = "github.com/xeipuuv/gojsonreference",
        sum = "h1:EzJWgHovont7NscjpAxXsDA8S8BMYve8Y5+7cuRE7R0=",
        version = "v0.0.0-20180127040603-bd5ef7bd5415",
    )
    go_repository(
        name = "com_github_xeipuuv_gojsonschema",
        build_file_proto_mode = "disable",
        importpath = "github.com/xeipuuv/gojsonschema",
        sum = "h1:LhYJRs+L4fBtjZUfuSZIKGeVu0QRy8e5Xi7D17UxZ74=",
        version = "v1.2.0",
    )
    go_repository(
        name = "com_github_xiang90_probing",
        build_file_proto_mode = "disable",
        importpath = "github.com/xiang90/probing",
        sum = "h1:eY9dn8+vbi4tKz5Qo6v2eYzo7kUS51QINcR5jNpbZS8=",
        version = "v0.0.0-20190116061207-43a291ad63a2",
    )
    go_repository(
        name = "com_github_yalp_jsonpath",
        build_file_proto_mode = "disable",
        importpath = "github.com/yalp/jsonpath",
        sum = "h1:6fRhSjgLCkTD3JnJxvaJ4Sj+TYblw757bqYgZaOq5ZY=",
        version = "v0.0.0-20180802001716-5cc68e5049a0",
    )
    go_repository(
        name = "com_github_youmark_pkcs8",
        build_file_proto_mode = "disable",
        importpath = "github.com/youmark/pkcs8",
        sum = "h1:fZHgsYlfvtyqToslyjUt3VOPF4J7aK/3MPcK7xp3PDk=",
        version = "v0.0.0-20201027041543-1326539a0a0a",
    )
    go_repository(
        name = "com_github_yudai_gojsondiff",
        build_file_proto_mode = "disable",
        importpath = "github.com/yudai/gojsondiff",
        sum = "h1:27cbfqXLVEJ1o8I6v3y9lg8Ydm53EKqHXAOMxEGlCOA=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_yudai_golcs",
        build_file_proto_mode = "disable",
        importpath = "github.com/yudai/golcs",
        sum = "h1:BHyfKlQyqbsFN5p3IfnEUduWvb9is428/nNb5L3U01M=",
        version = "v0.0.0-20170316035057-ecda9a501e82",
    )
    go_repository(
        name = "com_github_yudai_pp",
        build_file_proto_mode = "disable",
        importpath = "github.com/yudai/pp",
        sum = "h1:Q4//iY4pNF6yPLZIigmvcl7k/bPgrcTPIFIcmawg5bI=",
        version = "v2.0.1+incompatible",
    )
    go_repository(
        name = "com_github_yuin_goldmark",
        build_file_proto_mode = "disable",
        importpath = "github.com/yuin/goldmark",
        sum = "h1:fVcFKWvrslecOb/tg+Cc05dkeYx540o0FuFt3nUVDoE=",
        version = "v1.4.13",
    )
    go_repository(
        name = "com_github_yusufpapurcu_wmi",
        build_file_proto_mode = "disable",
        importpath = "github.com/yusufpapurcu/wmi",
        sum = "h1:KBNDSne4vP5mbSWnJbO+51IMOXJB67QiYCSBrubbPRg=",
        version = "v1.2.2",
    )
    go_repository(
        name = "com_github_zclconf_go_cty",
        build_file_proto_mode = "disable",
        importpath = "github.com/zclconf/go-cty",
        sum = "h1:PcupnljUm9EIvbgSHQnHhUr3fO6oFmkOrvs2BAFNXXY=",
        version = "v1.12.1",
    )
    go_repository(
        name = "com_github_zclconf_go_cty_debug",
        build_file_proto_mode = "disable",
        importpath = "github.com/zclconf/go-cty-debug",
        sum = "h1:FosyBZYxY34Wul7O/MSKey3txpPYyCqVO5ZyceuQJEI=",
        version = "v0.0.0-20191215020915-b22d67c1ba0b",
    )
    go_repository(
        name = "com_github_zenazn_goji",
        build_file_proto_mode = "disable",
        importpath = "github.com/zenazn/goji",
        sum = "h1:RSQQAbXGArQ0dIDEq+PI6WqN6if+5KHu6x2Cx/GXLTQ=",
        version = "v0.9.0",
    )
    go_repository(
        name = "com_google_cloud_go",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go",
        sum = "h1:DNtEKRBAAzeS4KyIory52wWHuClNaXJ5x1F7xa4q+5Y=",
        version = "v0.105.0",
    )
    go_repository(
        name = "com_google_cloud_go_accessapproval",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/accessapproval",
        sum = "h1:/nTivgnV/n1CaAeo+ekGexTYUsKEU9jUVkoY5359+3Q=",
        version = "v1.5.0",
    )
    go_repository(
        name = "com_google_cloud_go_accesscontextmanager",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/accesscontextmanager",
        sum = "h1:CFhNhU7pcD11cuDkQdrE6PQJgv0EXNKNv06jIzbLlCU=",
        version = "v1.4.0",
    )
    go_repository(
        name = "com_google_cloud_go_aiplatform",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/aiplatform",
        sum = "h1:DBi3Jk9XjCJ4pkkLM4NqKgj3ozUL1wq4l+d3/jTGXAI=",
        version = "v1.27.0",
    )
    go_repository(
        name = "com_google_cloud_go_analytics",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/analytics",
        sum = "h1:NKw6PpQi6V1O+KsjuTd+bhip9d0REYu4NevC45vtGp8=",
        version = "v0.12.0",
    )
    go_repository(
        name = "com_google_cloud_go_apigateway",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/apigateway",
        sum = "h1:IIoXKR7FKrEAQhMTz5hK2wiDz2WNFHS7eVr/L1lE/rM=",
        version = "v1.4.0",
    )
    go_repository(
        name = "com_google_cloud_go_apigeeconnect",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/apigeeconnect",
        sum = "h1:AONoTYJviyv1vS4IkvWzq69gEVdvHx35wKXc+e6wjZQ=",
        version = "v1.4.0",
    )
    go_repository(
        name = "com_google_cloud_go_appengine",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/appengine",
        sum = "h1:lmG+O5oaR9xNwaRBwE2XoMhwQHsHql5IoiGr1ptdDwU=",
        version = "v1.5.0",
    )
    go_repository(
        name = "com_google_cloud_go_area120",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/area120",
        sum = "h1:TCMhwWEWhCn8d44/Zs7UCICTWje9j3HuV6nVGMjdpYw=",
        version = "v0.6.0",
    )
    go_repository(
        name = "com_google_cloud_go_artifactregistry",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/artifactregistry",
        sum = "h1:3d0LRAU1K6vfqCahhl9fx2oGHcq+s5gftdix4v8Ibrc=",
        version = "v1.9.0",
    )
    go_repository(
        name = "com_google_cloud_go_asset",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/asset",
        sum = "h1:aCrlaLGJWTODJX4G56ZYzJefITKEWNfbjjtHSzWpxW0=",
        version = "v1.10.0",
    )
    go_repository(
        name = "com_google_cloud_go_assuredworkloads",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/assuredworkloads",
        sum = "h1:hhIdCOowsT1GG5eMCIA0OwK6USRuYTou/1ZeNxCSRtA=",
        version = "v1.9.0",
    )
    go_repository(
        name = "com_google_cloud_go_automl",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/automl",
        sum = "h1:BMioyXSbg7d7xLibn47cs0elW6RT780IUWr42W8rp2Q=",
        version = "v1.8.0",
    )
    go_repository(
        name = "com_google_cloud_go_baremetalsolution",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/baremetalsolution",
        sum = "h1:g9KO6SkakcYPcc/XjAzeuUrEOXlYPnMpuiaywYaGrmQ=",
        version = "v0.4.0",
    )
    go_repository(
        name = "com_google_cloud_go_batch",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/batch",
        sum = "h1:1jvEBY55OH4Sd2FxEXQfxGExFWov1A/IaRe+Z5Z71Fw=",
        version = "v0.4.0",
    )
    go_repository(
        name = "com_google_cloud_go_beyondcorp",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/beyondcorp",
        sum = "h1:w+4kThysgl0JiKshi2MKDCg2NZgOyqOI0wq2eBZyrzA=",
        version = "v0.3.0",
    )
    go_repository(
        name = "com_google_cloud_go_bigquery",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/bigquery",
        sum = "h1:Wi4dITi+cf9VYp4VH2T9O41w0kCW0uQTELq2Z6tukN0=",
        version = "v1.44.0",
    )
    go_repository(
        name = "com_google_cloud_go_billing",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/billing",
        sum = "h1:Xkii76HWELHwBtkQVZvqmSo9GTr0O+tIbRNnMcGdlg4=",
        version = "v1.7.0",
    )
    go_repository(
        name = "com_google_cloud_go_binaryauthorization",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/binaryauthorization",
        sum = "h1:pL70vXWn9TitQYXBWTK2abHl2JHLwkFRjYw6VflRqEA=",
        version = "v1.4.0",
    )
    go_repository(
        name = "com_google_cloud_go_certificatemanager",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/certificatemanager",
        sum = "h1:tzbR4UHBbgsewMWUD93JHi8EBi/gHBoSAcY1/sThFGk=",
        version = "v1.4.0",
    )
    go_repository(
        name = "com_google_cloud_go_channel",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/channel",
        sum = "h1:pNuUlZx0Jb0Ts9P312bmNMuH5IiFWIR4RUtLb70Ke5s=",
        version = "v1.9.0",
    )
    go_repository(
        name = "com_google_cloud_go_cloudbuild",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/cloudbuild",
        sum = "h1:TAAmCmAlOJ4uNBu6zwAjwhyl/7fLHHxIEazVhr3QBbQ=",
        version = "v1.4.0",
    )
    go_repository(
        name = "com_google_cloud_go_clouddms",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/clouddms",
        sum = "h1:UhzHIlgFfMr6luVYVNydw/pl9/U5kgtjCMJHnSvoVws=",
        version = "v1.4.0",
    )
    go_repository(
        name = "com_google_cloud_go_cloudtasks",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/cloudtasks",
        sum = "h1:faUiUgXjW8yVZ7XMnKHKm1WE4OldPBUWWfIRN/3z1dc=",
        version = "v1.8.0",
    )
    go_repository(
        name = "com_google_cloud_go_compute",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/compute",
        sum = "h1:7UGq3QknM33pw5xATlpzeoomNxsacIVvTqTTvbfajmE=",
        version = "v1.15.1",
    )
    go_repository(
        name = "com_google_cloud_go_compute_metadata",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/compute/metadata",
        sum = "h1:mg4jlk7mCAj6xXp9UJ4fjI9VUI5rubuGBW5aJ7UnBMY=",
        version = "v0.2.3",
    )

    go_repository(
        name = "com_google_cloud_go_contactcenterinsights",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/contactcenterinsights",
        sum = "h1:tTQLI/ZvguUf9Hv+36BkG2+/PeC8Ol1q4pBW+tgCx0A=",
        version = "v1.4.0",
    )
    go_repository(
        name = "com_google_cloud_go_container",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/container",
        sum = "h1:nbEK/59GyDRKKlo1SqpohY1TK8LmJ2XNcvS9Gyom2A0=",
        version = "v1.7.0",
    )
    go_repository(
        name = "com_google_cloud_go_containeranalysis",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/containeranalysis",
        sum = "h1:2824iym832ljKdVpCBnpqm5K94YT/uHTVhNF+dRTXPI=",
        version = "v0.6.0",
    )
    go_repository(
        name = "com_google_cloud_go_datacatalog",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/datacatalog",
        sum = "h1:6kZ4RIOW/uT7QWC5SfPfq/G8sYzr/v+UOmOAxy4Z1TE=",
        version = "v1.8.0",
    )
    go_repository(
        name = "com_google_cloud_go_dataflow",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/dataflow",
        sum = "h1:CW3541Fm7KPTyZjJdnX6NtaGXYFn5XbFC5UcjgALKvU=",
        version = "v0.7.0",
    )
    go_repository(
        name = "com_google_cloud_go_dataform",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/dataform",
        sum = "h1:vLwowLF2ZB5J5gqiZCzv076lDI/Rd7zYQQFu5XO1PSg=",
        version = "v0.5.0",
    )
    go_repository(
        name = "com_google_cloud_go_datafusion",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/datafusion",
        sum = "h1:j5m2hjWovTZDTQak4MJeXAR9yN7O+zMfULnjGw/OOLg=",
        version = "v1.5.0",
    )
    go_repository(
        name = "com_google_cloud_go_datalabeling",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/datalabeling",
        sum = "h1:dp8jOF21n/7jwgo/uuA0RN8hvLcKO4q6s/yvwevs2ZM=",
        version = "v0.6.0",
    )
    go_repository(
        name = "com_google_cloud_go_dataplex",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/dataplex",
        sum = "h1:cNxeA2DiWliQGi21kPRqnVeQ5xFhNoEjPRt1400Pm8Y=",
        version = "v1.4.0",
    )
    go_repository(
        name = "com_google_cloud_go_dataproc",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/dataproc",
        sum = "h1:gVOqNmElfa6n/ccG/QDlfurMWwrK3ezvy2b2eDoCmS0=",
        version = "v1.8.0",
    )
    go_repository(
        name = "com_google_cloud_go_dataqna",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/dataqna",
        sum = "h1:gx9jr41ytcA3dXkbbd409euEaWtofCVXYBvJz3iYm18=",
        version = "v0.6.0",
    )
    go_repository(
        name = "com_google_cloud_go_datastore",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/datastore",
        sum = "h1:4siQRf4zTiAVt/oeH4GureGkApgb2vtPQAtOmhpqQwE=",
        version = "v1.10.0",
    )
    go_repository(
        name = "com_google_cloud_go_datastream",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/datastream",
        sum = "h1:PgIgbhedBtYBU6POGXFMn2uSl9vpqubc3ewTNdcU8Mk=",
        version = "v1.5.0",
    )
    go_repository(
        name = "com_google_cloud_go_deploy",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/deploy",
        sum = "h1:kI6dxt8Ml0is/x7YZjLveTvR7YPzXAUD/8wQZ2nH5zA=",
        version = "v1.5.0",
    )
    go_repository(
        name = "com_google_cloud_go_dialogflow",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/dialogflow",
        sum = "h1:HYHVOkoxQ9bSfNIelSZYNAtUi4CeSrCnROyOsbOqPq8=",
        version = "v1.19.0",
    )
    go_repository(
        name = "com_google_cloud_go_dlp",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/dlp",
        sum = "h1:9I4BYeJSVKoSKgjr70fLdRDumqcUeVmHV4fd5f9LR6Y=",
        version = "v1.7.0",
    )
    go_repository(
        name = "com_google_cloud_go_documentai",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/documentai",
        sum = "h1:jfq09Fdjtnpnmt/MLyf6A3DM3ynb8B2na0K+vSXvpFM=",
        version = "v1.10.0",
    )
    go_repository(
        name = "com_google_cloud_go_domains",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/domains",
        sum = "h1:pu3JIgC1rswIqi5romW0JgNO6CTUydLYX8zyjiAvO1c=",
        version = "v0.7.0",
    )
    go_repository(
        name = "com_google_cloud_go_edgecontainer",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/edgecontainer",
        sum = "h1:hd6J2n5dBBRuAqnNUEsKWrp6XNPKsaxwwIyzOPZTokk=",
        version = "v0.2.0",
    )
    go_repository(
        name = "com_google_cloud_go_errorreporting",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/errorreporting",
        sum = "h1:kj1XEWMu8P0qlLhm3FwcaFsUvXChV/OraZwA70trRR0=",
        version = "v0.3.0",
    )
    go_repository(
        name = "com_google_cloud_go_essentialcontacts",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/essentialcontacts",
        sum = "h1:b6csrQXCHKQmfo9h3dG/pHyoEh+fQG1Yg78a53LAviY=",
        version = "v1.4.0",
    )
    go_repository(
        name = "com_google_cloud_go_eventarc",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/eventarc",
        sum = "h1:AgCqrmMMIcel5WWKkzz5EkCUKC3Rl5LNMMYsS+LvsI0=",
        version = "v1.8.0",
    )
    go_repository(
        name = "com_google_cloud_go_filestore",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/filestore",
        sum = "h1:yjKOpzvqtDmL5AXbKttLc8j0hL20kuC1qPdy5HPcxp0=",
        version = "v1.4.0",
    )
    go_repository(
        name = "com_google_cloud_go_firestore",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/firestore",
        sum = "h1:IBlRyxgGySXu5VuW0RgGFlTtLukSnNkpDiEOMkQkmpA=",
        version = "v1.9.0",
    )
    go_repository(
        name = "com_google_cloud_go_functions",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/functions",
        sum = "h1:35tgv1fQOtvKqH/uxJMzX3w6usneJ0zXpsFr9KAVhNE=",
        version = "v1.9.0",
    )
    go_repository(
        name = "com_google_cloud_go_gaming",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/gaming",
        sum = "h1:97OAEQtDazAJD7yh/kvQdSCQuTKdR0O+qWAJBZJ4xiA=",
        version = "v1.8.0",
    )
    go_repository(
        name = "com_google_cloud_go_gkebackup",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/gkebackup",
        sum = "h1:4K+jiv4ocqt1niN8q5Imd8imRoXBHTrdnJVt/uFFxF4=",
        version = "v0.3.0",
    )
    go_repository(
        name = "com_google_cloud_go_gkeconnect",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/gkeconnect",
        sum = "h1:zAcvDa04tTnGdu6TEZewaLN2tdMtUOJJ7fEceULjguA=",
        version = "v0.6.0",
    )
    go_repository(
        name = "com_google_cloud_go_gkehub",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/gkehub",
        sum = "h1:JTcTaYQRGsVm+qkah7WzHb6e9sf1C0laYdRPn9aN+vg=",
        version = "v0.10.0",
    )
    go_repository(
        name = "com_google_cloud_go_gkemulticloud",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/gkemulticloud",
        sum = "h1:8F1NhJj8ucNj7lK51UZMtAjSWTgP1zO18XF6vkfiPPU=",
        version = "v0.4.0",
    )
    go_repository(
        name = "com_google_cloud_go_gsuiteaddons",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/gsuiteaddons",
        sum = "h1:TGT2oGmO5q3VH6SjcrlgPUWI0njhYv4kywLm6jag0to=",
        version = "v1.4.0",
    )
    go_repository(
        name = "com_google_cloud_go_iam",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/iam",
        sum = "h1:E2osAkZzxI/+8pZcxVLcDtAQx/u+hZXVryUaYQ5O0Kk=",
        version = "v0.8.0",
    )
    go_repository(
        name = "com_google_cloud_go_iap",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/iap",
        sum = "h1:BGEXovwejOCt1zDk8hXq0bOhhRu9haXKWXXXp2B4wBM=",
        version = "v1.5.0",
    )
    go_repository(
        name = "com_google_cloud_go_ids",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/ids",
        sum = "h1:LncHK4HHucb5Du310X8XH9/ICtMwZ2PCfK0ScjWiJoY=",
        version = "v1.2.0",
    )
    go_repository(
        name = "com_google_cloud_go_iot",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/iot",
        sum = "h1:Y9+oZT9jD4GUZzORXTU45XsnQrhxmDT+TFbPil6pRVQ=",
        version = "v1.4.0",
    )
    go_repository(
        name = "com_google_cloud_go_kms",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/kms",
        sum = "h1:OWRZzrPmOZUzurjI2FBGtgY2mB1WaJkqhw6oIwSj0Yg=",
        version = "v1.6.0",
    )
    go_repository(
        name = "com_google_cloud_go_language",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/language",
        sum = "h1:3Wa+IUMamL4JH3Zd3cDZUHpwyqplTACt6UZKRD2eCL4=",
        version = "v1.8.0",
    )
    go_repository(
        name = "com_google_cloud_go_lifesciences",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/lifesciences",
        sum = "h1:tIqhivE2LMVYkX0BLgG7xL64oNpDaFFI7teunglt1tI=",
        version = "v0.6.0",
    )
    go_repository(
        name = "com_google_cloud_go_logging",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/logging",
        sum = "h1:ZBsZK+JG+oCDT+vaxwqF2egKNRjz8soXiS6Xv79benI=",
        version = "v1.6.1",
    )
    go_repository(
        name = "com_google_cloud_go_longrunning",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/longrunning",
        sum = "h1:NjljC+FYPV3uh5/OwWT6pVU+doBqMg2x/rZlE+CamDs=",
        version = "v0.3.0",
    )
    go_repository(
        name = "com_google_cloud_go_managedidentities",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/managedidentities",
        sum = "h1:3Kdajn6X25yWQFhFCErmKSYTSvkEd3chJROny//F1A0=",
        version = "v1.4.0",
    )
    go_repository(
        name = "com_google_cloud_go_maps",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/maps",
        sum = "h1:kLReRbclTgJefw2fcCbdLPLhPj0U6UUWN10ldG8sdOU=",
        version = "v0.1.0",
    )
    go_repository(
        name = "com_google_cloud_go_mediatranslation",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/mediatranslation",
        sum = "h1:qAJzpxmEX+SeND10Y/4868L5wfZpo4Y3BIEnIieP4dk=",
        version = "v0.6.0",
    )
    go_repository(
        name = "com_google_cloud_go_memcache",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/memcache",
        sum = "h1:yLxUzJkZVSH2kPaHut7k+7sbIBFpvSh1LW9qjM2JDjA=",
        version = "v1.7.0",
    )
    go_repository(
        name = "com_google_cloud_go_metastore",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/metastore",
        sum = "h1:3KcShzqWdqxrDEXIBWpYJpOOrgpDj+HlBi07Grot49Y=",
        version = "v1.8.0",
    )
    go_repository(
        name = "com_google_cloud_go_monitoring",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/monitoring",
        sum = "h1:c9riaGSPQ4dUKWB+M1Fl0N+iLxstMbCktdEwYSPGDvA=",
        version = "v1.8.0",
    )
    go_repository(
        name = "com_google_cloud_go_networkconnectivity",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/networkconnectivity",
        sum = "h1:BVdIKaI68bihnXGdCVL89Jsg9kq2kg+II30fjVqo62E=",
        version = "v1.7.0",
    )
    go_repository(
        name = "com_google_cloud_go_networkmanagement",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/networkmanagement",
        sum = "h1:mDHA3CDW00imTvC5RW6aMGsD1bH+FtKwZm/52BxaiMg=",
        version = "v1.5.0",
    )
    go_repository(
        name = "com_google_cloud_go_networksecurity",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/networksecurity",
        sum = "h1:qDEX/3sipg9dS5JYsAY+YvgTjPR63cozzAWop8oZS94=",
        version = "v0.6.0",
    )
    go_repository(
        name = "com_google_cloud_go_notebooks",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/notebooks",
        sum = "h1:AC8RPjNvel3ExgXjO1YOAz+teg9+j+89TNxa7pIZfww=",
        version = "v1.5.0",
    )
    go_repository(
        name = "com_google_cloud_go_optimization",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/optimization",
        sum = "h1:7PxOq9VTT7TMib/6dMoWpMvWS2E4dJEvtYzjvBreaec=",
        version = "v1.2.0",
    )
    go_repository(
        name = "com_google_cloud_go_orchestration",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/orchestration",
        sum = "h1:39d6tqvNjd/wsSub1Bn4cEmrYcet5Ur6xpaN+SxOxtY=",
        version = "v1.4.0",
    )
    go_repository(
        name = "com_google_cloud_go_orgpolicy",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/orgpolicy",
        sum = "h1:erF5PHqDZb6FeFrUHiYj2JK2BMhsk8CyAg4V4amJ3rE=",
        version = "v1.5.0",
    )
    go_repository(
        name = "com_google_cloud_go_osconfig",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/osconfig",
        sum = "h1:NO0RouqCOM7M2S85Eal6urMSSipWwHU8evzwS+siqUI=",
        version = "v1.10.0",
    )
    go_repository(
        name = "com_google_cloud_go_oslogin",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/oslogin",
        sum = "h1:pKGDPfeZHDybtw48WsnVLjoIPMi9Kw62kUE5TXCLCN4=",
        version = "v1.7.0",
    )
    go_repository(
        name = "com_google_cloud_go_phishingprotection",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/phishingprotection",
        sum = "h1:OrwHLSRSZyaiOt3tnY33dsKSedxbMzsXvqB21okItNQ=",
        version = "v0.6.0",
    )
    go_repository(
        name = "com_google_cloud_go_policytroubleshooter",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/policytroubleshooter",
        sum = "h1:NQklJuOUoz1BPP+Epjw81COx7IISWslkZubz/1i0UN8=",
        version = "v1.4.0",
    )
    go_repository(
        name = "com_google_cloud_go_privatecatalog",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/privatecatalog",
        sum = "h1:Vz86uiHCtNGm1DeC32HeG2VXmOq5JRYA3VRPf8ZEcSg=",
        version = "v0.6.0",
    )
    go_repository(
        name = "com_google_cloud_go_pubsub",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/pubsub",
        sum = "h1:q+J/Nfr6Qx4RQeu3rJcnN48SNC0qzlYzSeqkPq93VHs=",
        version = "v1.27.1",
    )
    go_repository(
        name = "com_google_cloud_go_pubsublite",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/pubsublite",
        sum = "h1:iqrD8vp3giTb7hI1q4TQQGj77cj8zzgmMPsTZtLnprM=",
        version = "v1.5.0",
    )
    go_repository(
        name = "com_google_cloud_go_recaptchaenterprise_v2",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/recaptchaenterprise/v2",
        sum = "h1:UqzFfb/WvhwXGDF1eQtdHLrmni+iByZXY4h3w9Kdyv8=",
        version = "v2.5.0",
    )
    go_repository(
        name = "com_google_cloud_go_recommendationengine",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/recommendationengine",
        sum = "h1:6w+WxPf2LmUEqX0YyvfCoYb8aBYOcbIV25Vg6R0FLGw=",
        version = "v0.6.0",
    )
    go_repository(
        name = "com_google_cloud_go_recommender",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/recommender",
        sum = "h1:9kMZQGeYfcOD/RtZfcNKGKtoex3DdoB4zRgYU/WaIwE=",
        version = "v1.8.0",
    )
    go_repository(
        name = "com_google_cloud_go_redis",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/redis",
        sum = "h1:/zTwwBKIAD2DEWTrXZp8WD9yD/gntReF/HkPssVYd0U=",
        version = "v1.10.0",
    )
    go_repository(
        name = "com_google_cloud_go_resourcemanager",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/resourcemanager",
        sum = "h1:NDao6CHMwEZIaNsdWy+tuvHaavNeGP06o1tgrR0kLvU=",
        version = "v1.4.0",
    )
    go_repository(
        name = "com_google_cloud_go_resourcesettings",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/resourcesettings",
        sum = "h1:eTzOwB13WrfF0kuzG2ZXCfB3TLunSHBur4s+HFU6uSM=",
        version = "v1.4.0",
    )
    go_repository(
        name = "com_google_cloud_go_retail",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/retail",
        sum = "h1:N9fa//ecFUOEPsW/6mJHfcapPV0wBSwIUwpVZB7MQ3o=",
        version = "v1.11.0",
    )
    go_repository(
        name = "com_google_cloud_go_run",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/run",
        sum = "h1:AWPuzU7Xtaj3Jf+QarDWIs6AJ5hM1VFQ+F6Q+VZ6OT4=",
        version = "v0.3.0",
    )
    go_repository(
        name = "com_google_cloud_go_scheduler",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/scheduler",
        sum = "h1:K/mxOewgHGeKuATUJNGylT75Mhtjmx1TOkKukATqMT8=",
        version = "v1.7.0",
    )
    go_repository(
        name = "com_google_cloud_go_secretmanager",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/secretmanager",
        sum = "h1:xE6uXljAC1kCR8iadt9+/blg1fvSbmenlsDN4fT9gqw=",
        version = "v1.9.0",
    )
    go_repository(
        name = "com_google_cloud_go_security",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/security",
        sum = "h1:KSKzzJMyUoMRQzcz7azIgqAUqxo7rmQ5rYvimMhikqg=",
        version = "v1.10.0",
    )
    go_repository(
        name = "com_google_cloud_go_securitycenter",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/securitycenter",
        sum = "h1:QTVtk/Reqnx2bVIZtJKm1+mpfmwRwymmNvlaFez7fQY=",
        version = "v1.16.0",
    )
    go_repository(
        name = "com_google_cloud_go_servicecontrol",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/servicecontrol",
        sum = "h1:ImIzbOu6y4jL6ob65I++QzvqgFaoAKgHOG+RU9/c4y8=",
        version = "v1.5.0",
    )
    go_repository(
        name = "com_google_cloud_go_servicedirectory",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/servicedirectory",
        sum = "h1:f7M8IMcVzO3T425AqlZbP3yLzeipsBHtRza8vVFYMhQ=",
        version = "v1.7.0",
    )
    go_repository(
        name = "com_google_cloud_go_servicemanagement",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/servicemanagement",
        sum = "h1:TpkCO5M7dhKSy1bKUD9o/sSEW/U1Gtx7opA1fsiMx0c=",
        version = "v1.5.0",
    )
    go_repository(
        name = "com_google_cloud_go_serviceusage",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/serviceusage",
        sum = "h1:b0EwJxPJLpavSljMQh0RcdHsUrr5DQ+Nelt/3BAs5ro=",
        version = "v1.4.0",
    )
    go_repository(
        name = "com_google_cloud_go_shell",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/shell",
        sum = "h1:b1LFhFBgKsG252inyhtmsUUZwchqSz3WTvAIf3JFo4g=",
        version = "v1.4.0",
    )
    go_repository(
        name = "com_google_cloud_go_spanner",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/spanner",
        sum = "h1:NvdTpRwf7DTegbfFdPjAWyD7bOVu0VeMqcvR9aCQCAc=",
        version = "v1.41.0",
    )
    go_repository(
        name = "com_google_cloud_go_speech",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/speech",
        sum = "h1:yK0ocnFH4Wsf0cMdUyndJQ/hPv02oTJOxzi6AgpBy4s=",
        version = "v1.9.0",
    )
    go_repository(
        name = "com_google_cloud_go_storage",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/storage",
        sum = "h1:6RRlFMv1omScs6iq2hfE3IvgE+l6RfJPampq8UZc5TU=",
        version = "v1.14.0",
    )
    go_repository(
        name = "com_google_cloud_go_storagetransfer",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/storagetransfer",
        sum = "h1:fUe3OydbbvHcAYp07xY+2UpH4AermGbmnm7qdEj3tGE=",
        version = "v1.6.0",
    )
    go_repository(
        name = "com_google_cloud_go_talent",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/talent",
        sum = "h1:MrekAGxLqAeAol4Sc0allOVqUGO8j+Iim8NMvpiD7tM=",
        version = "v1.4.0",
    )
    go_repository(
        name = "com_google_cloud_go_texttospeech",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/texttospeech",
        sum = "h1:ccPiHgTewxgyAeCWgQWvZvrLmbfQSFABTMAfrSPLPyY=",
        version = "v1.5.0",
    )
    go_repository(
        name = "com_google_cloud_go_tpu",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/tpu",
        sum = "h1:ztIdKoma1Xob2qm6QwNh4Xi9/e7N3IfvtwG5AcNsj1g=",
        version = "v1.4.0",
    )
    go_repository(
        name = "com_google_cloud_go_trace",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/trace",
        sum = "h1:qO9eLn2esajC9sxpqp1YKX37nXC3L4BfGnPS0Cx9dYo=",
        version = "v1.4.0",
    )
    go_repository(
        name = "com_google_cloud_go_translate",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/translate",
        sum = "h1:AOYOH3MspzJ/bH1YXzB+xTE8fMpn3mwhLjugwGXvMPI=",
        version = "v1.4.0",
    )
    go_repository(
        name = "com_google_cloud_go_video",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/video",
        sum = "h1:ttlvO4J5c1VGq6FkHqWPD/aH6PfdxujHt+muTJlW1Zk=",
        version = "v1.9.0",
    )
    go_repository(
        name = "com_google_cloud_go_videointelligence",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/videointelligence",
        sum = "h1:RPFgVVXbI2b5vnrciZjtsUgpNKVtHO/WIyXUhEfuMhA=",
        version = "v1.9.0",
    )
    go_repository(
        name = "com_google_cloud_go_vision_v2",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/vision/v2",
        sum = "h1:TQHxRqvLMi19azwm3qYuDbEzZWmiKJNTpGbkNsfRCik=",
        version = "v2.5.0",
    )
    go_repository(
        name = "com_google_cloud_go_vmmigration",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/vmmigration",
        sum = "h1:A2Tl2ZmwMRpvEmhV2ibISY85fmQR+Y5w9a0PlRz5P3s=",
        version = "v1.3.0",
    )
    go_repository(
        name = "com_google_cloud_go_vmwareengine",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/vmwareengine",
        sum = "h1:JMPZaOT/gIUxVlTqSl/QQ32Y2k+r0stNeM1NSqhVP9o=",
        version = "v0.1.0",
    )
    go_repository(
        name = "com_google_cloud_go_vpcaccess",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/vpcaccess",
        sum = "h1:woHXXtnW8b9gLFdWO9HLPalAddBQ9V4LT+1vjKwR3W8=",
        version = "v1.5.0",
    )
    go_repository(
        name = "com_google_cloud_go_webrisk",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/webrisk",
        sum = "h1:ypSnpGlJnZSXbN9a13PDmAYvVekBLnGKxQ3Q9SMwnYY=",
        version = "v1.7.0",
    )
    go_repository(
        name = "com_google_cloud_go_websecurityscanner",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/websecurityscanner",
        sum = "h1:y7yIFg/h/mO+5Y5aCOtVAnpGUOgqCH5rXQ2Oc8Oq2+g=",
        version = "v1.4.0",
    )
    go_repository(
        name = "com_google_cloud_go_workflows",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/workflows",
        sum = "h1:7Chpin9p50NTU8Tb7qk+I11U/IwVXmDhEoSsdccvInE=",
        version = "v1.9.0",
    )
    go_repository(
        name = "com_shuralyov_dmitri_gpu_mtl",
        build_file_proto_mode = "disable",
        importpath = "dmitri.shuralyov.com/gpu/mtl",
        sum = "h1:VpgP7xuJadIUuKccphEpTJnWhS2jkQyMt6Y7pJCD7fY=",
        version = "v0.0.0-20190408044501-666a987793e9",
    )
    go_repository(
        name = "com_sourcegraph_sourcegraph_appdash",
        build_file_proto_mode = "disable",
        importpath = "sourcegraph.com/sourcegraph/appdash",
        sum = "h1:ucqkfpjg9WzSUubAO62csmucvxl4/JeW3F4I4909XkM=",
        version = "v0.0.0-20190731080439-ebfcffb1b5c0",
    )
    go_repository(
        name = "in_gopkg_alecthomas_kingpin_v2",
        build_file_proto_mode = "disable",
        importpath = "gopkg.in/alecthomas/kingpin.v2",
        sum = "h1:jMFz6MfLP0/4fUyZle81rXUoxOBFi19VUFKVDOQfozc=",
        version = "v2.2.6",
    )
    go_repository(
        name = "in_gopkg_check_v1",
        build_file_proto_mode = "disable",
        importpath = "gopkg.in/check.v1",
        sum = "h1:Hei/4ADfdWqJk1ZMxUNpqntNwaWcugrBjAiHlqqRiVk=",
        version = "v1.0.0-20201130134442-10cb98267c6c",
    )
    go_repository(
        name = "in_gopkg_cheggaaa_pb_v1",
        build_file_proto_mode = "disable",
        importpath = "gopkg.in/cheggaaa/pb.v1",
        sum = "h1:Ev7yu1/f6+d+b3pi5vPdRPc6nNtP1umSfcWiEfRqv6I=",
        version = "v1.0.25",
    )
    go_repository(
        name = "in_gopkg_errgo_v2",
        build_file_proto_mode = "disable",
        importpath = "gopkg.in/errgo.v2",
        sum = "h1:0vLT13EuvQ0hNvakwLuFZ/jYrLp5F3kcWHXdRggjCE8=",
        version = "v2.1.0",
    )
    go_repository(
        name = "in_gopkg_fsnotify_v1",
        build_file_proto_mode = "disable",
        importpath = "gopkg.in/fsnotify.v1",
        sum = "h1:xOHLXZwVvI9hhs+cLKq5+I5onOuwQLhQwiu63xxlHs4=",
        version = "v1.4.7",
    )
    go_repository(
        name = "in_gopkg_gcfg_v1",
        build_file_proto_mode = "disable",
        importpath = "gopkg.in/gcfg.v1",
        sum = "h1:m8OOJ4ccYHnx2f4gQwpno8nAX5OGOh7RLaaz0pj3Ogs=",
        version = "v1.2.3",
    )
    go_repository(
        name = "in_gopkg_go_playground_assert_v1",
        build_file_proto_mode = "disable",
        importpath = "gopkg.in/go-playground/assert.v1",
        sum = "h1:xoYuJVE7KT85PYWrN730RguIQO0ePzVRfFMXadIrXTM=",
        version = "v1.2.1",
    )
    go_repository(
        name = "in_gopkg_go_playground_validator_v8",
        build_file_proto_mode = "disable",
        importpath = "gopkg.in/go-playground/validator.v8",
        sum = "h1:lFB4DoMU6B626w8ny76MV7VX6W2VHct2GVOI3xgiMrQ=",
        version = "v8.18.2",
    )

    go_repository(
        name = "in_gopkg_inconshreveable_log15_v2",
        build_file_proto_mode = "disable",
        importpath = "gopkg.in/inconshreveable/log15.v2",
        sum = "h1:RlWgLqCMMIYYEVcAR5MDsuHlVkaIPDAF+5Dehzg8L5A=",
        version = "v2.0.0-20180818164646-67afb5ed74ec",
    )
    go_repository(
        name = "in_gopkg_inf_v0",
        build_file_proto_mode = "disable",
        importpath = "gopkg.in/inf.v0",
        sum = "h1:73M5CoZyi3ZLMOyDlQh031Cx6N9NDJ2Vvfl76EDAgDc=",
        version = "v0.9.1",
    )

    go_repository(
        name = "in_gopkg_ini_v1",
        build_file_proto_mode = "disable",
        importpath = "gopkg.in/ini.v1",
        sum = "h1:Dgnx+6+nfE+IfzjUEISNeydPJh9AXNNsWbGP9KzCsOA=",
        version = "v1.67.0",
    )
    go_repository(
        name = "in_gopkg_natefinch_lumberjack_v2",
        build_file_proto_mode = "disable",
        importpath = "gopkg.in/natefinch/lumberjack.v2",
        sum = "h1:1Lc07Kr7qY4U2YPouBjpCLxpiyxIVoxqXgkXLknAOE8=",
        version = "v2.0.0",
    )

    go_repository(
        name = "in_gopkg_resty_v1",
        build_file_proto_mode = "disable",
        importpath = "gopkg.in/resty.v1",
        sum = "h1:CuXP0Pjfw9rOuY6EP+UvtNvt5DSqHpIxILZKT/quCZI=",
        version = "v1.12.0",
    )
    go_repository(
        name = "in_gopkg_tomb_v1",
        build_file_proto_mode = "disable",
        importpath = "gopkg.in/tomb.v1",
        sum = "h1:uRGJdciOHaEIrze2W8Q3AKkepLTh2hOroT7a+7czfdQ=",
        version = "v1.0.0-20141024135613-dd632973f1e7",
    )
    go_repository(
        name = "in_gopkg_warnings_v0",
        build_file_proto_mode = "disable",
        importpath = "gopkg.in/warnings.v0",
        sum = "h1:wFXVbFY8DY5/xOe1ECiWdKCzZlxgshcYVNkBHstARME=",
        version = "v0.1.2",
    )
    go_repository(
        name = "in_gopkg_yaml_v2",
        build_file_proto_mode = "disable",
        importpath = "gopkg.in/yaml.v2",
        sum = "h1:D8xgwECY7CYvx+Y2n4sBz93Jn9JRvxdiyyo8CTfuKaY=",
        version = "v2.4.0",
    )
    go_repository(
        name = "in_gopkg_yaml_v3",
        build_file_proto_mode = "disable",
        importpath = "gopkg.in/yaml.v3",
        sum = "h1:fxVm/GzAzEWqLHuvctI91KS9hhNmmWOoWu0XTYJS7CA=",
        version = "v3.0.1",
    )
    go_repository(
        name = "io_ariga_atlas",
        build_file_proto_mode = "disable",
        importpath = "ariga.io/atlas",
        sum = "h1:mbsFtavDqGdYwdDpP50LGOOZ2hgyGoJcZeOpbgKMyu4=",
        version = "v0.9.1-0.20230119145809-92243f7c55cb",
    )
    go_repository(
        name = "io_entgo_contrib",
        build_file_proto_mode = "disable",
        importpath = "entgo.io/contrib",
        sum = "h1:wY85TgRp3j5ix/SZ9IE6Ob5lObHFmVUYH0ZFw1D5Hzc=",
        version = "v0.3.5",
    )

    go_repository(
        name = "io_entgo_ent",
        build_file_proto_mode = "disable",
        importpath = "entgo.io/ent",
        sum = "h1:M/M0QL1CYCUSdqGRXUrXhFYSDRJPsOOrr+RLEej/gyQ=",
        version = "v0.11.8",
    )
    go_repository(
        name = "io_etcd_go_bbolt",
        build_file_proto_mode = "disable",
        importpath = "go.etcd.io/bbolt",
        sum = "h1:/ecaJf0sk1l4l6V4awd65v2C3ILy7MSj+s/x1ADCIMU=",
        version = "v1.3.6",
    )
    go_repository(
        name = "io_etcd_go_etcd",
        build_file_proto_mode = "disable",
        importpath = "go.etcd.io/etcd",
        sum = "h1:VcrIfasaLFkyjk6KNlXQSzO+B0fZcnECiDrKJsfxka0=",
        version = "v0.0.0-20191023171146-3cf2f69b5738",
    )
    go_repository(
        name = "io_etcd_go_etcd_api_v3",
        build_file_proto_mode = "disable",
        importpath = "go.etcd.io/etcd/api/v3",
        sum = "h1:BX4JIbQ7hl7+jL+g+2j5UAr0o1bctCm6/Ct+ArBGkf0=",
        version = "v3.5.5",
    )
    go_repository(
        name = "io_etcd_go_etcd_client_pkg_v3",
        build_file_proto_mode = "disable",
        importpath = "go.etcd.io/etcd/client/pkg/v3",
        sum = "h1:9S0JUVvmrVl7wCF39iTQthdaaNIiAaQbmK75ogO6GU8=",
        version = "v3.5.5",
    )
    go_repository(
        name = "io_etcd_go_etcd_client_v2",
        build_file_proto_mode = "disable",
        importpath = "go.etcd.io/etcd/client/v2",
        sum = "h1:DktRP60//JJpnPC0VBymAN/7V71GHMdjDCBt4ZPXDjI=",
        version = "v2.305.5",
    )
    go_repository(
        name = "io_etcd_go_etcd_client_v3",
        build_file_proto_mode = "disable",
        importpath = "go.etcd.io/etcd/client/v3",
        sum = "h1:q++2WTJbUgpQu4B6hCuT7VkdwaTP7Qz6Daak3WzbrlI=",
        version = "v3.5.5",
    )
    go_repository(
        name = "io_etcd_go_etcd_pkg_v3",
        build_file_proto_mode = "disable",
        importpath = "go.etcd.io/etcd/pkg/v3",
        sum = "h1:ntrg6vvKRW26JRmHTE0iNlDgYK6JX3hg/4cD62X0ixk=",
        version = "v3.5.0",
    )
    go_repository(
        name = "io_etcd_go_etcd_raft_v3",
        build_file_proto_mode = "disable",
        importpath = "go.etcd.io/etcd/raft/v3",
        sum = "h1:kw2TmO3yFTgE+F0mdKkG7xMxkit2duBDa2Hu6D/HMlw=",
        version = "v3.5.0",
    )
    go_repository(
        name = "io_etcd_go_etcd_server_v3",
        build_file_proto_mode = "disable",
        importpath = "go.etcd.io/etcd/server/v3",
        sum = "h1:jk8D/lwGEDlQU9kZXUFMSANkE22Sg5+mW27ip8xcF9E=",
        version = "v3.5.0",
    )
    go_repository(
        name = "io_k8s_api",
        build_file_proto_mode = "disable",
        importpath = "k8s.io/api",
        sum = "h1:f+SWYiPd/GsiWwVRz+NbFyCgvv75Pk9NK6dlkZgpCRQ=",
        version = "v0.26.1",
    )
    go_repository(
        name = "io_k8s_apimachinery",
        build_file_proto_mode = "disable",
        importpath = "k8s.io/apimachinery",
        sum = "h1:8EZ/eGJL+hY/MYCNwhmDzVqq2lPl3N3Bo8rvweJwXUQ=",
        version = "v0.26.1",
    )
    go_repository(
        name = "io_k8s_client_go",
        build_file_proto_mode = "disable",
        importpath = "k8s.io/client-go",
        sum = "h1:87CXzYJnAMGaa/IDDfRdhTzxk/wzGZ+/HUQpqgVSZXU=",
        version = "v0.26.1",
    )
    go_repository(
        name = "io_k8s_gengo",
        build_file_proto_mode = "disable",
        importpath = "k8s.io/gengo",
        sum = "h1:GohjlNKauSai7gN4wsJkeZ3WAJx4Sh+oT/b5IYn5suA=",
        version = "v0.0.0-20210813121822-485abfe95c7c",
    )
    go_repository(
        name = "io_k8s_klog",
        build_file_proto_mode = "disable",
        importpath = "k8s.io/klog",
        sum = "h1:Pt+yjF5aB1xDSVbau4VsWe+dQNzA0qv1LlXdC2dF6Q8=",
        version = "v1.0.0",
    )
    go_repository(
        name = "io_k8s_klog_v2",
        build_file_proto_mode = "disable",
        importpath = "k8s.io/klog/v2",
        sum = "h1:atnLQ121W371wYYFawwYx1aEY2eUfs4l3J72wtgAwV4=",
        version = "v2.80.1",
    )
    go_repository(
        name = "io_k8s_kube_openapi",
        build_file_proto_mode = "disable",
        importpath = "k8s.io/kube-openapi",
        sum = "h1:+70TFaan3hfJzs+7VK2o+OGxg8HsuBr/5f6tVAjDu6E=",
        version = "v0.0.0-20221012153701-172d655c2280",
    )
    go_repository(
        name = "io_k8s_sigs_json",
        build_file_proto_mode = "disable",
        importpath = "sigs.k8s.io/json",
        sum = "h1:iXTIw73aPyC+oRdyqqvVJuloN1p0AC/kzH07hu3NE+k=",
        version = "v0.0.0-20220713155537-f223a00ba0e2",
    )
    go_repository(
        name = "io_k8s_sigs_structured_merge_diff",
        build_file_proto_mode = "disable",
        importpath = "sigs.k8s.io/structured-merge-diff",
        sum = "h1:4Z09Hglb792X0kfOBBJUPFEyvVfQWrYT/l8h5EKA6JQ=",
        version = "v0.0.0-20190525122527-15d366b2352e",
    )
    go_repository(
        name = "io_k8s_sigs_structured_merge_diff_v4",
        build_file_proto_mode = "disable",
        importpath = "sigs.k8s.io/structured-merge-diff/v4",
        sum = "h1:PRbqxJClWWYMNV1dhaG4NsibJbArud9kFxnAMREiWFE=",
        version = "v4.2.3",
    )

    go_repository(
        name = "io_k8s_sigs_yaml",
        build_file_proto_mode = "disable",
        importpath = "sigs.k8s.io/yaml",
        sum = "h1:a2VclLzOGrwOHDiV8EfBGhvjHvP46CtW5j6POvhYGGo=",
        version = "v1.3.0",
    )
    go_repository(
        name = "io_k8s_utils",
        build_file_proto_mode = "disable",
        importpath = "k8s.io/utils",
        sum = "h1:0Smp/HP1OH4Rvhe+4B8nWGERtlqAGSftbSbbmm45oFs=",
        version = "v0.0.0-20221107191617-1a15be271d1d",
    )

    go_repository(
        name = "io_opencensus_go",
        build_file_proto_mode = "disable",
        importpath = "go.opencensus.io",
        sum = "h1:y73uSU6J157QMP2kn2r30vwW1A2W2WFwSCGnAVxeaD0=",
        version = "v0.24.0",
    )
    go_repository(
        name = "io_opentelemetry_go_contrib",
        build_file_proto_mode = "disable",
        importpath = "go.opentelemetry.io/contrib",
        sum = "h1:ubFQUn0VCZ0gPwIoJfBJVpeBlyRMxu8Mm/huKWYd9p0=",
        version = "v0.20.0",
    )
    go_repository(
        name = "io_opentelemetry_go_contrib_instrumentation_google_golang_org_grpc_otelgrpc",
        build_file_proto_mode = "disable",
        importpath = "go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc",
        sum = "h1:sO4WKdPAudZGKPcpZT4MJn6JaDmpyLrMPDGGyA1SttE=",
        version = "v0.20.0",
    )

    go_repository(
        name = "io_opentelemetry_go_otel",
        build_file_proto_mode = "disable",
        importpath = "go.opentelemetry.io/otel",
        sum = "h1:1ZAKnNQKwBBxFtww/GwxNUyTf0AxkZzrukO8MeXqe4Y=",
        version = "v1.13.0",
    )
    go_repository(
        name = "io_opentelemetry_go_otel_exporters_jaeger",
        build_file_proto_mode = "disable",
        importpath = "go.opentelemetry.io/otel/exporters/jaeger",
        sum = "h1:VAMoGujbVV8Q0JNM/cEbhzUIWWBxnEqH45HP9iBKN04=",
        version = "v1.13.0",
    )
    go_repository(
        name = "io_opentelemetry_go_otel_exporters_otlp",
        build_file_proto_mode = "disable",
        importpath = "go.opentelemetry.io/otel/exporters/otlp",
        sum = "h1:PTNgq9MRmQqqJY0REVbZFvwkYOA85vbdQU/nVfxDyqg=",
        version = "v0.20.0",
    )
    go_repository(
        name = "io_opentelemetry_go_otel_exporters_zipkin",
        build_file_proto_mode = "disable",
        importpath = "go.opentelemetry.io/otel/exporters/zipkin",
        sum = "h1:RqPV1VhJjrx28qOKYFPj3Mso56uaBovur3GZehF9y9s=",
        version = "v1.13.0",
    )
    go_repository(
        name = "io_opentelemetry_go_otel_metric",
        build_file_proto_mode = "disable",
        importpath = "go.opentelemetry.io/otel/metric",
        sum = "h1:4kzhXFP+btKm4jwxpjIqjs41A7MakRFUS86bqLHTIw8=",
        version = "v0.20.0",
    )

    go_repository(
        name = "io_opentelemetry_go_otel_oteltest",
        build_file_proto_mode = "disable",
        importpath = "go.opentelemetry.io/otel/oteltest",
        sum = "h1:G685iP3XiskCwk/z0eIabL55XUl2gk0cljhGk9sB0Yk=",
        version = "v1.0.0-RC1",
    )
    go_repository(
        name = "io_opentelemetry_go_otel_sdk",
        build_file_proto_mode = "disable",
        importpath = "go.opentelemetry.io/otel/sdk",
        sum = "h1:BHib5g8MvdqS65yo2vV1s6Le42Hm6rrw08qU6yz5JaM=",
        version = "v1.13.0",
    )
    go_repository(
        name = "io_opentelemetry_go_otel_sdk_export_metric",
        build_file_proto_mode = "disable",
        importpath = "go.opentelemetry.io/otel/sdk/export/metric",
        sum = "h1:c5VRjxCXdQlx1HjzwGdQHzZaVI82b5EbBgOu2ljD92g=",
        version = "v0.20.0",
    )
    go_repository(
        name = "io_opentelemetry_go_otel_sdk_metric",
        build_file_proto_mode = "disable",
        importpath = "go.opentelemetry.io/otel/sdk/metric",
        sum = "h1:7ao1wpzHRVKf0OQ7GIxiQJA6X7DLX9o14gmVon7mMK8=",
        version = "v0.20.0",
    )

    go_repository(
        name = "io_opentelemetry_go_otel_trace",
        build_file_proto_mode = "disable",
        importpath = "go.opentelemetry.io/otel/trace",
        sum = "h1:CBgRZ6ntv+Amuj1jDsMhZtlAPT6gbyIRdaIzFhfBSdY=",
        version = "v1.13.0",
    )
    go_repository(
        name = "io_opentelemetry_go_proto_otlp",
        build_file_proto_mode = "disable",
        importpath = "go.opentelemetry.io/proto/otlp",
        sum = "h1:rwOQPCuKAKmwGKq2aVNnYIibI6wnV7EvzgfTCzcdGg8=",
        version = "v0.7.0",
    )
    go_repository(
        name = "io_rsc_binaryregexp",
        build_file_proto_mode = "disable",
        importpath = "rsc.io/binaryregexp",
        sum = "h1:HfqmD5MEmC0zvwBuF187nq9mdnXjXsSivRiXN7SmRkE=",
        version = "v0.2.0",
    )
    go_repository(
        name = "io_rsc_quote_v3",
        build_file_proto_mode = "disable",
        importpath = "rsc.io/quote/v3",
        sum = "h1:9JKUTTIUgS6kzR9mK1YuGKv6Nl+DijDNIc0ghT58FaY=",
        version = "v3.1.0",
    )
    go_repository(
        name = "io_rsc_sampler",
        build_file_proto_mode = "disable",
        importpath = "rsc.io/sampler",
        sum = "h1:7uVkIFmeBqHfdjD+gZwtXXI+RODJ2Wc4O7MPEh/QiW4=",
        version = "v1.3.0",
    )
    go_repository(
        name = "org_golang_google_api",
        build_file_proto_mode = "disable",
        importpath = "google.golang.org/api",
        sum = "h1:JxJl2qQ85fRMPNvlZY/enexbxpCjLwGhZUtgfGeQ51I=",
        version = "v0.102.0",
    )
    go_repository(
        name = "org_golang_google_appengine",
        build_file_proto_mode = "disable",
        importpath = "google.golang.org/appengine",
        sum = "h1:FZR1q0exgwxzPzp/aF+VccGrSfxfPpkBqjIIEq3ru6c=",
        version = "v1.6.7",
    )
    go_repository(
        name = "org_golang_google_genproto",
        build_file_proto_mode = "disable",
        importpath = "google.golang.org/genproto",
        sum = "h1:BWUVssLB0HVOSY78gIdvk1dTVYtT1y8SBWtPYuTJ/6w=",
        version = "v0.0.0-20230110181048-76db0878b65f",
    )
    go_repository(
        name = "org_golang_google_grpc",
        build_file_proto_mode = "disable",
        importpath = "google.golang.org/grpc",
        sum = "h1:LAv2ds7cmFV/XTS3XG1NneeENYrXGmorPxsBbptIjNc=",
        version = "v1.53.0",
    )
    go_repository(
        name = "org_golang_google_grpc_cmd_protoc_gen_go_grpc",
        build_file_proto_mode = "disable",
        importpath = "google.golang.org/grpc/cmd/protoc-gen-go-grpc",
        sum = "h1:M1YKkFIboKNieVO5DLUEVzQfGwJD30Nv2jfUgzb5UcE=",
        version = "v1.1.0",
    )

    go_repository(
        name = "org_golang_google_protobuf",
        build_file_proto_mode = "disable",
        importpath = "google.golang.org/protobuf",
        sum = "h1:d0NfwRgPtno5B1Wa6L2DAG+KivqkdutMf1UhdNx175w=",
        version = "v1.28.1",
    )
    go_repository(
        name = "org_golang_x_crypto",
        build_file_proto_mode = "disable",
        importpath = "golang.org/x/crypto",
        sum = "h1:qfktjS5LUO+fFKeJXZ+ikTRijMmljikvG68fpMMruSc=",
        version = "v0.6.0",
    )
    go_repository(
        name = "org_golang_x_exp",
        build_file_proto_mode = "disable",
        importpath = "golang.org/x/exp",
        sum = "h1:m9O6OTJ627iFnN2JIWfdqlZCzneRO6EEBsHXI25P8ws=",
        version = "v0.0.0-20221230185412-738e83a70c30",
    )
    go_repository(
        name = "org_golang_x_image",
        build_file_proto_mode = "disable",
        importpath = "golang.org/x/image",
        sum = "h1:/DcQ0w3VHKCC5p0/P2B0JpAZ9Z++V2KOo2fyU89CXBQ=",
        version = "v0.2.0",
    )
    go_repository(
        name = "org_golang_x_lint",
        build_file_proto_mode = "disable",
        importpath = "golang.org/x/lint",
        sum = "h1:VLliZ0d+/avPrXXH+OakdXhpJuEoBZuwh1m2j7U6Iug=",
        version = "v0.0.0-20210508222113-6edffad5e616",
    )
    go_repository(
        name = "org_golang_x_mobile",
        build_file_proto_mode = "disable",
        importpath = "golang.org/x/mobile",
        sum = "h1:4+4C/Iv2U4fMZBiMCc98MG1In4gJY5YRhtpDNeDeHWs=",
        version = "v0.0.0-20190719004257-d2bd2a29d028",
    )
    go_repository(
        name = "org_golang_x_mod",
        build_file_proto_mode = "disable",
        importpath = "golang.org/x/mod",
        sum = "h1:LapD9S96VoQRhi/GrNTqeBJFrUjs5UHCAtTlgwA5oZA=",
        version = "v0.7.0",
    )
    go_repository(
        name = "org_golang_x_net",
        build_file_proto_mode = "disable",
        importpath = "golang.org/x/net",
        sum = "h1:rJrUqqhjsgNp7KqAIc25s9pZnjU7TUcSY7HcVZjdn1g=",
        version = "v0.7.0",
    )
    go_repository(
        name = "org_golang_x_oauth2",
        build_file_proto_mode = "disable",
        importpath = "golang.org/x/oauth2",
        sum = "h1:NF0gk8LVPg1Ml7SSbGyySuoxdsXitj7TvgvuRxIMc/M=",
        version = "v0.4.0",
    )
    go_repository(
        name = "org_golang_x_sync",
        build_file_proto_mode = "disable",
        importpath = "golang.org/x/sync",
        sum = "h1:wsuoTGHzEhffawBOhz5CYhcrV4IdKZbEyZjBMuTp12o=",
        version = "v0.1.0",
    )
    go_repository(
        name = "org_golang_x_sys",
        build_file_proto_mode = "disable",
        importpath = "golang.org/x/sys",
        sum = "h1:MUK/U/4lj1t1oPg0HfuXDN/Z1wv31ZJ/YcPiGccS4DU=",
        version = "v0.5.0",
    )
    go_repository(
        name = "org_golang_x_term",
        build_file_proto_mode = "disable",
        importpath = "golang.org/x/term",
        sum = "h1:n2a8QNdAb0sZNpU9R1ALUXBbY+w51fCQDN+7EdxNBsY=",
        version = "v0.5.0",
    )
    go_repository(
        name = "org_golang_x_text",
        build_file_proto_mode = "disable",
        importpath = "golang.org/x/text",
        sum = "h1:4BRB4x83lYWy72KwLD/qYDuTu7q9PjSagHvijDw7cLo=",
        version = "v0.7.0",
    )
    go_repository(
        name = "org_golang_x_time",
        build_file_proto_mode = "disable",
        importpath = "golang.org/x/time",
        sum = "h1:rg5rLMjNzMS1RkNLzCG38eapWhnYLFYXDXj2gOlr8j4=",
        version = "v0.3.0",
    )
    go_repository(
        name = "org_golang_x_tools",
        build_file_proto_mode = "disable",
        importpath = "golang.org/x/tools",
        sum = "h1:7mTAgkunk3fr4GAloyyCasadO6h9zSsQZbwvcaIciV4=",
        version = "v0.4.0",
    )
    go_repository(
        name = "org_golang_x_xerrors",
        build_file_proto_mode = "disable",
        importpath = "golang.org/x/xerrors",
        sum = "h1:H2TDz8ibqkAF6YGhCdN3jS9O0/s90v0rJh3X/OLHEUk=",
        version = "v0.0.0-20220907171357-04be3eba64a2",
    )
    go_repository(
        name = "org_mongodb_go_mongo_driver",
        build_file_proto_mode = "disable",
        importpath = "go.mongodb.org/mongo-driver",
        sum = "h1:QP0znIRTuL0jf1oBQoAoM0C6ZJfBK4kx0Uumtv1A7w8=",
        version = "v1.11.1",
    )
    go_repository(
        name = "org_uber_go_atomic",
        build_file_proto_mode = "disable",
        importpath = "go.uber.org/atomic",
        sum = "h1:9qC72Qh0+3MqyJbAn8YU5xVq1frD8bn3JtD2oXtafVQ=",
        version = "v1.10.0",
    )
    go_repository(
        name = "org_uber_go_automaxprocs",
        build_file_proto_mode = "disable",
        importpath = "go.uber.org/automaxprocs",
        sum = "h1:e1YG66Lrk73dn4qhg8WFSvhF0JuFQF0ERIp4rpuV8Qk=",
        version = "v1.5.1",
    )

    go_repository(
        name = "org_uber_go_goleak",
        build_file_proto_mode = "disable",
        importpath = "go.uber.org/goleak",
        sum = "h1:wy28qYRKZgnJTxGxvye5/wgWr1EKjmUDGYox5mGlRlI=",
        version = "v1.1.11",
    )
    go_repository(
        name = "org_uber_go_multierr",
        build_file_proto_mode = "disable",
        importpath = "go.uber.org/multierr",
        sum = "h1:7fIwc/ZtS0q++VgcfqFDxSBZVv/Xo49/SYnDFupUwlI=",
        version = "v1.9.0",
    )
    go_repository(
        name = "org_uber_go_ratelimit",
        build_file_proto_mode = "disable",
        importpath = "go.uber.org/ratelimit",
        sum = "h1:U2AruXqeTb4Eh9sYQSTrMhH8Cb7M0Ian2ibBOnBcnAw=",
        version = "v0.1.0",
    )

    go_repository(
        name = "org_uber_go_tools",
        build_file_proto_mode = "disable",
        importpath = "go.uber.org/tools",
        sum = "h1:0mgffUl7nfd+FpvXMVz4IDEaUSmT1ysygQC7qYo7sG4=",
        version = "v0.0.0-20190618225709-2cfd321de3ee",
    )
    go_repository(
        name = "org_uber_go_zap",
        build_file_proto_mode = "disable",
        importpath = "go.uber.org/zap",
        sum = "h1:FiJd5l1UOLj0wCgbSE0rwwXHzEdAZS6hiiSnxJN/D60=",
        version = "v1.24.0",
    )
