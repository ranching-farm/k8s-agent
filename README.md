<img src="https://ranching.farm/images/logo.svg" alt="ranching.farm logo"/>

# [ranching.farm](https://ranching.farm) Kubernetes Agent

[![Build Status](https://img.shields.io/github/workflow/status/ranching-farm/k8s-agent/CI?style=flat-square)](https://github.com/ranching-farm/k8s-agent/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/ranching-farm/k8s-agent?style=flat-square)](https://goreportcard.com/report/github.com/ranching-farm/k8s-agent)
[![License](https://img.shields.io/github/license/ranching-farm/k8s-agent?style=flat-square)](https://github.com/ranching-farm/k8s-agent/blob/main/LICENSE)
[![Release](https://img.shields.io/github/v/release/ranching-farm/k8s-agent?style=flat-square)](https://github.com/ranching-farm/k8s-agent/releases)

## What is ranching.farm?

ranching.farm is a Kubernetes management platform that leverages AI to simplify cluster operations and troubleshooting. It's designed for DevOps engineers and SREs who want to streamline their Kubernetes workflows and get intelligent assistance for common tasks. It can also be used as a training platform and e-learning resource for beginners and intermediates alike.

We offer two ways of connecting your k8s deployment to our AI:

1) Use our k8s-agent (**this repository**) to deploy it straight into your kubernetes cluster.
2) Use our [kubectl-addon](https://github.com/ranching-farm/kubectl-addon) to run it locally on your machine using your kubectl configuration.

## What does ranching.farm do?

Once you are connected to ranching.farm through either the k8s-agent or the kubectl-addon, you gain access to a powerful AI-assisted Kubernetes management experience. Here's what you can do:

- **AI-Powered Chat**: Interact with your cluster using natural language through either OpenAI's latest ChatGPT model or Anthropic's Claude. Choose the AI that best fits your needs.

- **Kubernetes Command Execution**: The AI can utilize `kubectl`, `helm`, and `kustomize` to perform a wide range of operations on your cluster.

- **Cluster Debugging**: Get assistance in debugging complex issues, including DNS problems within your cluster.

- **Intelligent Analysis**: Receive insights and recommendations based on your cluster's state and configuration.

- **Learning and Training**: Use the platform as an interactive learning tool to improve your Kubernetes skills.

ranching.farm acts as an intelligent assistant, helping you manage, troubleshoot, and optimize your Kubernetes environment more efficiently.

This Kubernetes agent allows you to connect your Kubernetes cluster to the [ranching.farm](https://ranching.farm) platform. Once connected, you can interact with your cluster through our AI-powered interface.

## 🚀 Quick Start

1. Visit [ranching.farm](https://ranching.farm) and add your cluster to your account.
2. Copy the kubectl command provided on the website.
3. Run the command in your terminal to deploy the agent to your cluster:

```bash
kubectl apply -f <URL_provided_on_website>
```

## 🌟 Features

- **Easy Setup**: Deploy the agent to your cluster with a single command
- **Lightweight**: Minimal resource footprint
- **Secure**: Runs with limited permissions within your cluster
- **AI-Powered Assistance**: Access to ranching.farm's AI Kubernetes management features
- **Cross-Cluster Support**: Connect multiple clusters to your ranching.farm account

## 🛠 Installation

1. Log in to your [ranching.farm](https://ranching.farm) account.
2. Navigate to the "Add Cluster" section.
3. Follow the on-screen instructions to generate a unique kubectl command for your cluster.
4. Run the provided kubectl command in your terminal to deploy the agent.

## 📄 License

This project is licensed under the BSD 3-Clause License - see the [LICENSE](LICENSE) file for details.

## 🙋‍♀️ Support

- 🐛 [Issue Tracker](https://github.com/ranching-farm/k8s-agent/issues)

---

<p align="center">Made with ❤️ by the ranching.farm team</p>