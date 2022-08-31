# GoAlert

GoAlert provides on-call scheduling, automated escalations and notifications (like SMS or voice calls) to automatically engage the right person, the right way, and at the right time.

![image](https://user-images.githubusercontent.com/11381794/187750677-23f22ad2-1e91-4aec-a32d-76cd993d6c59.png)

![image](https://user-images.githubusercontent.com/11381794/187750834-2cd01f28-17bd-497b-93b4-6d2bee4940c9.png)

## Installation

GoAlert is distributed as a single binary with release notes available from the [GitHub Releases](https://github.com/target/goalert/releases) page.

See our [Getting Started Guide](./docs/getting-started.md) for running GoAlert in a production environment.

### Quick Start

```bash
# podman
podman run -it --rm -p 8081:8081 goalert/demo

# docker
docker run -it --rm -p 8081:8081 goalert/demo
```

GoAlert will be running at [localhost:8081](http://localhost:8081). You can login with `admin/admin123`.

## Contributing

If you'd like to contribute to GoAlert, please see our [Contributing Guidelines](./CONTRIBUTING.md) and the [Development Setup Guide](./docs/development-setup.md).

Please also see our [Code of Conduct](./CODE_OF_CONDUCT.md).

## Contact Us

If you need help or have a question, the `#goalert` Slack channel is available on [gophers.slack.com](https://gophers.slack.com/messages/goalert/).

To access Gophers Slack and the `#goalert` channel, you will need an invitation. You request one through the automated process here: https://invite.slack.golangbridge.org/

- Vote on existing [Feature Requests](https://github.com/target/goalert/issues?q=is%3Aopen+label%3Aenhancement+sort%3Areactions-%2B1-desc) or submit [a new one](https://github.com/target/goalert/issues/new)
- File a [bug report](https://github.com/target/goalert/issues)
- Report security issues to security@goalert.me

## License

GoAlert is licensed under the [Apache License, Version 2.0](./LICENSE.md).
