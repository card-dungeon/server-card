pipeline {
    // install golang 1.14 on Jenkins node
    agent any
    tools {
        go 'go 1.18.3'
    }
    environment {
        GOPATH = "${JENKINS_HOME}/jobs/${JOB_NAME}/builds/${BUILD_ID}"
    }
    stages {
        // stage("unit-test") {
        //     steps {
        //         echo 'UNIT TEST EXECUTION STARTED'
        //         sh 'make unit-tests'
        //     }
        // }
        // stage("functional-test") {
        //     steps {
        //         echo 'FUNCTIONAL TEST EXECUTION STARTED'
        //         sh 'make functional-tests'
        //     }
        // }
        stage("build") {
            steps {
                echo 'BUILD EXECUTION STARTED'
                sh 'docker build . -t gmae199boy/server-card'
            }
        }
        // stage('deliver') {
        //     agent any
        //     steps {
        //         withCredentials([usernamePassword(credentialsId: 'dockerhub_id', passwordVariable: 'dockerhubPassword', usernameVariable: 'dockerhubUser')]) {
        //         sh "docker login -u ${env.dockerhubUser} -p ${env.dockerhubPassword}"
        //         sh 'docker push gmae199boy/server-card'
        //         }
        //     }
        // }
        stage('Docker') {         
            environment {
                // Extract the username and password of our credentials into "DOCKER_CREDENTIALS_USR" and "DOCKER_CREDENTIALS_PSW".
                // (NOTE 1: DOCKER_CREDENTIALS will be set to "your_username:your_password".)
                // The new variables will always be YOUR_VARIABLE_NAME + _USR and _PSW.
                // (NOTE 2: You can't print credentials in the pipeline for security reasons.)
                DOCKER_CREDENTIALS = credentials('dockerhun_id')
            }

            steps {                           
                // Use a scripted pipeline.
                script {
                    node {
                        def app

                        stage('Clone repository') {
                            checkout scm
                        }

                        stage('Build image') {                            
                            app = docker.build("${env.DOCKER_CREDENTIALS_USR}/server-card")
                        }

                        stage('Push image') {  
                            // Use the Credential ID of the Docker Hub Credentials we added to Jenkins.
                            docker.withRegistry('https://registry.hub.docker.com', 'dockerhub_id') {                                
                                // Push image and tag it with our build number for versioning purposes.
                                // app.push("${env.BUILD_NUMBER}")                      
                                app.push(0.0.1)                      

                                // Push the same image and tag it as the latest version (appears at the top of our version list).
                                app.push("latest")
                            }
                        }              
                    }                 
                }
            }
        }

    }
}