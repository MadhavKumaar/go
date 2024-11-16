pipeline {
    agent any  // Use any available agent to run the pipeline

    environment {
        // Define environment variables
        APP_NAME = "MyApp"
        DEPLOY_ENV = "staging"
    }

    stages {
        stage('Checkout') {
            steps {
                // Checkout code from the repository
                git 'https://github.com/your-repository/your-project.git'
            }
        }

        stage('Build') {
            steps {
                script {
                    // Run your build commands (e.g., Maven, Gradle, npm)
                    echo "Building the application..."
                    sh 'mvn clean install'  // Example for Maven-based projects
                }
            }
        }

        stage('Test') {
            steps {
                script {
                    // Run your tests (e.g., unit tests)
                    echo "Running tests..."
                    sh 'mvn test'  // Example for running tests with Maven
                }
            }
        }

        stage('Static Code Analysis') {
            steps {
                script {
                    // Run static analysis tools (e.g., SonarQube, Checkstyle)
                    echo "Running static code analysis..."
                    sh 'mvn sonar:sonar'  // Example for SonarQube analysis
                }
            }
        }

        stage('Deploy') {
            when {
                branch 'main'  // Only deploy if the branch is 'main'
            }
            steps {
                script {
                    // Deploy to your environment (e.g., Kubernetes, AWS, Heroku)
                    echo "Deploying to $DEPLOY_ENV environment..."
                    sh './deploy.sh'  // Custom deploy script (could be any deployment method)
                }
            }
        }
    }

    post {
        success {
            echo 'Pipeline completed successfully!'
        }
        failure {
            echo 'Pipeline failed. Check the logs for errors.'
        }
        always {
            echo 'This will always run, regardless of success or failure.'
        }
    }
}
