pipeline {
    agent any

    environment {
        DOCKER_IMAGE_NAME = 'go-fiber-app'
        DOCKER_TAG = 'latest'
        REGISTRY = 'registry.sintek.com'  // Ganti dengan URL registry jika menggunakan selain Docker Hub
        DOCKER_CREDENTIALS = 'docker-registry-credentials' // Ganti dengan credentials ID di Jenkins
    }

    stages {
        stage('Checkout') {
            steps {
                // Checkout kode dari repository Git
                checkout scm
            }
        }

        stage('Build Docker Image') {
            steps {
                script {
                    // Membangun Docker image dari Dockerfile
                    sh "docker build -t ${REGISTRY}/${DOCKER_IMAGE_NAME}:${DOCKER_TAG} ."
                }
            }
        }

        stage('Login to Docker Registry') {
            steps {
                script {
                    // Login ke Docker registry (menggunakan credentials Jenkins)
                    docker.withCredentials([usernamePassword(credentialsId: DOCKER_CREDENTIALS, usernameVariable: 'DOCKER_USERNAME', passwordVariable: 'DOCKER_PASSWORD')]) {
                        sh "echo $DOCKER_PASSWORD | docker login -u $DOCKER_USERNAME --password-stdin ${REGISTRY}"
                    }
                }
            }
        }

        //stage('Push Docker Image') {
        //    steps {
        //        script {
        //            // Push Docker image ke registry
        //            sh "docker push https://${REGISTRY}/${DOCKER_IMAGE_NAME}:${DOCKER_TAG}"
        //        }
        //    }
        //}

        stage('Clean Up') {
            steps {
                // Menghapus Docker image lokal untuk menghemat ruang disk
                script {
                    sh "docker rmi ${REGISTRY}/${DOCKER_IMAGE_NAME}:${DOCKER_TAG}"
                }
            }
        }
    }

    post {
        always {
            // Menjaga Jenkins build tetap bersih, menghapus Docker images yang tidak diperlukan
            cleanWs()
        }
    }
}
