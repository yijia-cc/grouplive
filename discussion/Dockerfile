# Build
FROM maven:3.8.1-jdk-11-slim AS builder

WORKDIR /app

COPY pom.xml ./

RUN mvn dependency:go-offline

COPY src ./src

RUN mvn -B clean package

# Production
FROM maven:3.8.1-jdk-11-slim AS production

WORKDIR /app

COPY --from=builder /app/target/*.jar app.jar

ENTRYPOINT ["java", "-jar", "app.jar"]
