<a name="readme-top"></a>
<!--
*** Thanks for checking out the Best-README-Template. If you have a suggestion
*** that would make this better, please fork the repo and create a pull request
*** or simply open an issue with the tag "enhancement".
*** Don't forget to give the project a star!
*** Thanks again! Now go create something AMAZING! :D
-->



<!-- CONTACT ME -->
You can connect with me through LinkedIn using the link the following link: [![LinkedIn][linkedin-shield]][linkedin-url]

<!-- TABLE OF CONTENTS -->
<details>
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a>
      <ul>
        <li><a href="#built-with">Built With</a></li>
      </ul>
    </li>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#prerequisites">Prerequisites</a></li>
        <li><a href="#installation">Installation</a></li>
      </ul>
    </li>
    <li>
      <a href="#usage">Usage</a>
      <ul>
        <li><a href="#prometheus">Prometheus</a></li>
        <li><a href="#grafana">Grafana</a></li>
      </ul>
    </li>
    <li><a href="#contributing">Contributing</a></li>
    <!-- <li><a href="#license">License</a></li> -->
    <li><a href="#contact">Contact</a></li>
    <li><a href="#acknowledgments">Acknowledgments</a></li>
  </ol>
</details>



<!-- ABOUT THE PROJECT -->
## About The Project

This is a project to share my approach to building a basic go website using multiple technologies that can be seen in the <a href="#built-with">Built With</a> section. The motivation behind this project is to display the learning process of new technologies involved in building this project as well as its CI/CD pipeline. The successful functional end result should be as displayed here if you follow the <a href="#functional-run">Functional Run</a>:

![go-app-results-page][go-app-results-page]

However if you are just interested in the serving of the webapp in the kubernetes cluster you should follow the <a href="#basic-run">Basic Run</a> which would make the successful end result looking like:

![go-app-landing-page][go-app-landing-page]


<p align="right">(<a href="#readme-top">back to top</a>)</p>



### Built With

* [![Go][Go]][Go-url]
* [![Docker][Docker]][Docker-url]
* [![Kubernetes][Kubernetes]][Kubernetes-url]
* [![Github][Github]][Github-url]
* [![ArgoCD][ArgoCD]][ArgoCD-url]

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- GETTING STARTED -->
## Getting Started

Here are the sequential steps to get the project up and running!

### Prerequisites

* Ensure that you have docker installed on your machine. To install docker please follow the instructions of installation on the <a href="https://docker.com/">Docker Official Website</a>
according to your operating system.

* Ensure that docker is installed correctly by typing the following command in your shell and getting a response that looks like this
  ```sh
  ~$ docker --version
  Docker version XX.XX.XX ...
  ```
  >At the time of initiating the project version 24.0.2 of Docker was used.

* Ensure that you have minikube installed on your machine to simulate a kubernetes cluster locally. To install minikube please follow the instructions of installation on the <a href="https://minikube.sigs.k8s.io/docs/start/">Minikube Offical Installation Page</a> according to your operating system.

* Ensure that minikube is installed correctly by typing the following command in your shell and getting a response that looks like this
  ```sh
  ~$ minikube version
  minikube version: vX.XX.X
  commit: XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
  ```
  >At the time of initiating the project version 1.30.1 of Minikube was used.

* *OPTIONAL*: Install ArgoCD CLI to try and get familiar with controlling the ArgoCD pipeline from the terminal. Use the instructions on the <a href="https://argo-cd.readthedocs.io/en/stable/getting_started/">ArgoCD Installation Page</a>. You can check if the installation was successful by running the follwing command in your shell and getting a response that looks like this
  ```sh
  ~$ argocd version
    argocd: vX.X.X+XXXXXXX
        BuildDate: XXXX-XX-XXXXX:XX:XXX
        GitCommit: XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
        GitTreeState: XXXXX
        GoVersion: goX.XX.XX
        Compiler: XX
        Platform: XXXXXXX/XXXXX
    argocd-server: vX.X.X+XXXXXXX.XXXXX
  ```
  >At the time of initiating the project version 2.7.7 of argocd cli was used.

### Installation

1. Using Minikube with Docker as the driver, start a kubernetes cluster with 3 worker nodes running kubernetes version 1.27.3 by running the following command
   ```sh
   ~$ minikube start --driver=docker --kubernetes-version v1.27.3 --nodes 3
   ```
   > This process might take some time if it is the first start up of minikube.

2. Install ArgoCD in the kubernetes cluster by running the following commands
   ```sh
   ~$ kubectl create namespace argocd
   ~$ kubectl apply -n argocd -f https://raw.githubusercontent.com/argoproj/argo-cd/stable/manifests/core-install.yaml
   ```
   >These commands are taken from <a href="https://argo-cd.readthedocs.io/en/stable/getting_started/">ArgoCD Installation Page</a>. Make sure that the offical page takes priority over the commands in above.

3. Login to the ArgoCD UI.
    * Run the following command to serve the ArgoCD service on your local machine on a given port.
        ```sh
        ~$ kubectl port-forward -n argocd svc/argocd-server YOUR_DESIRED_PORT:443
        ```

    * Get the initial admin password of ArgoCD to be able to access the ArgoCD service through the UI or CLI through the following command.
        ```sh
        ~$ kubectl get secret -n argocd argocd-initial-admin-secret -o yaml
        apiVersion: v1
        data:
            password: XXXXXXXXXXXXXXXXXXXXXXXX
        kind: Secret
        metadata:
            creationTimestamp: "XXXX-XX-XXXXX:XX:XXX"
            name: argocd-initial-admin-secret
            namespace: argocd
            resourceVersion: "XXXX"
            uid: XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX
        type: Opaque
        ```
        __IMPORTANT: The password is base64 encoded so you need to decode the password before using it__
        * Base64 Decode using a UNIX shell like sh, bash, zsh:
            ```sh
            ~$ echo "VEhJU19JU19USEVfUEFTU1dPUkQ=" | base64 --decode
            THIS_IS_THE_PASSWORD
            ```
        * Base64 Decode using powershell for a Windows Operating System:
            ```powershell
            ~$ echo $([Text.Encoding]::Utf8.GetString([Convert]::FromBase64String("VEhJU19JU19USEVfUEFTU1dPUkQ=")))
            THIS_IS_THE_PASSWORD
            ```

    * Open your Browser and go the link `https://localhost:YOUR_DESIRED_PORT`<br>
      NOTE: On the browser it will show that it does not trust the ArgoCD website however this only due to the SSL certificate being self-signed and not signed by some CA (Certificate Authority) that your browser is aware of and trusts. Please procceed to the link without worry for security issues, then you will be able to see the ArgoCD UI.

4. Create an application.yaml file for ArgoCD with the content of the file in `gitops_kubernetes_manifests/application.yaml` located in the repository.

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- USAGE EXAMPLES -->
## Usage

In this section you can see how easy it is to get this kubernetes cluster up and running.

### Basic Run
1. Apply the created `application.yaml` file using the following command
    ```sh
    ~$ kubectl apply -f path/to/application.yaml
    ```
    The application should pop up almost immediatly on the ArgoCD UI open in the browser. Click on the application that popped up and you should see something similar to
    ![argocd-preview][argocd-preview]

2. Run the following command to serve the Go webapp
    ```sh
    ~$ kubectl port-forward -n development svc/go-webapp-service YOUR_DESIRED_GO_WEBAPP_PORT:8080
    ```
    > __NOTE__: make sure that YOUR_DESIRED_GO_WEBAPP_PORT is different from the ArgoCD UI port.

3. Open your Browser and go the link `https://localhost:YOUR_DESIRED_GO_WEBAPP_PORT`
    ![go-app-landing-page][go-app-landing-page]

    __IMPORTANT__: This makes you see that the cluster is up and running serving the Go webapp however the webapp is not functional for it needs a proper API key that you get directly from <a href="https://newsapi.org/">newsapi</a>. Check the next section *Functional Run*.

### Functional Run
1. create your own __PRIVATE REPO__ with the content in <a href="https://github.com/FaisalAl-Rayes/go_k8_argocd">this repo</a> so you can actually freely supply the API key into version control.

2. Give access to ArgoCD of your __PRIVATE REPO__ following the <a href="https://argo-cd.readthedocs.io/en/stable/user-guide/private-repositories/">ArgoCD Private Repositories Official Page</a> 

3. Change the content of the `gitops_kubernetes_manifests\overlays\development\secrets\news_api_key.secret.example` file in your __PRIVATE REPO__ to the API key you get from <a href="https://newsapi.org/">newsapi</a>

4. Change the `application.yaml` file's `spec.source.repoURL` to your __PRIVATE REPO__ url and then apply it with the following command
    ```sh
    ~$ kubectl apply -f path/to/application.yaml
    ```

5. Open your Browser and go the link `https://localhost:YOUR_DESIRED_GO_WEBAPP_PORT` and you will have the functioning version of this basic Go webapp
    ![go-app-results-page][go-app-results-page]

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- CONTRIBUTING -->
## Contributing

Contributions are what make the open source community such an amazing place to learn, inspire, and create. Any contributions you make are **greatly appreciated**.

If you have a suggestion that would make this better, please fork the repo and create a pull request. You can also simply open an issue with the tag "enhancement".
Don't forget to give the project a star! Thanks again!

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- LICENSE 
## License

Distributed under the MIT License. See `LICENSE.txt` for more information.

<p align="right">(<a href="#readme-top">back to top</a>)</p>
-->


<!-- CONTACT -->
## Contact
Connect with me on 

[![LinkedIn][linkedin-shield]][linkedin-url]

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- ACKNOWLEDGMENTS -->
## Acknowledgments

* [Best README Template](https://github.com/othneildrew/Best-README-Template/)
* [Original News Demo Project](https://github.com/Freshman-tech/news-demo-starter-files)

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- MARKDOWN LINKS & IMAGES -->
<!-- https://markdownguide.org/basic-syntax/#reference-style-links -->

[linkedin-shield]: https://img.shields.io/badge/linkedin-0769AD?style=for-the-badge&logo=linkedin&logoColor=white
[linkedin-url]: https://linkedin.com/in/faisalalrayyess

[argocd-preview]: readme_images/argocd_application.png
[go-app-landing-page]: readme_images/go_app_landing_page.png
[go-app-results-page]: readme_images/go_app_results_page.png

[Go]: https://img.shields.io/badge/go-306998?style=for-the-badge&logo=go&logoColor=white
[Go-url]: https://go.dev/

[Docker]: https://img.shields.io/badge/docker-0769AD?style=for-the-badge&logo=docker&logoColor=white
[Docker-url]: https://docker.com/

[kubernetes]: https://img.shields.io/badge/kubernetes-F5F5F5?style=for-the-badge&logo=kubernetes&logoColor=3970e4
[kubernetes-url]: https://kubernetes.io/

[Github]: https://img.shields.io/badge/github-1a1a1a?style=for-the-badge&logo=github&logoColor=F5F5F5
[Github-url]: https://github.com/

[ArgoCD]: https://img.shields.io/badge/argo-F5F5F5?style=for-the-badge&logo=argo&logoColor=orange
[ArgoCD-url]: https://argoproj.github.io/cd/