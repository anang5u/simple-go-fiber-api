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
                    // Build Docker image dari Dockerfile
                    sh "docker build -t ${REGISTRY}/${DOCKER_IMAGE_NAME}:${DOCKER_TAG} ."
                }
            }
        }

        stage('Login to Docker Registry') {
            steps {
                script {
                    // Menggunakan dengan kredensial untuk login ke Docker registry
                    withCredentials([usernamePassword(credentialsId: DOCKER_CREDENTIALS, usernameVariable: 'DOCKER_USERNAME', passwordVariable: 'DOCKER_PASSWORD')]) {
                        // Melakukan login menggunakan username dan password dari kredensial
                        sh "echo $DOCKER_PASSWORD | docker login -u $DOCKER_USERNAME --password-stdin ${REGISTRY}"
                    }
                }
            }
        }

        stage('Push Docker Image') {
            steps {
                script {
                    // Push Docker image ke registry
                    sh "docker push ${REGISTRY}/${DOCKER_IMAGE_NAME}:${DOCKER_TAG}"
                }
            }
        }

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
            // menggunakan plugin Workspace Cleanup Plugin
            //cleanWs()

            // Membersihkan file yang ada di workspace secara manual jika plugin cleanWs tidak ada
            sh 'rm -rf *'
        }
    }
}
