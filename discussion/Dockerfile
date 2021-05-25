# Build
FROM maven:3.8.1-jdk-11-slim AS builder

WORKDIR /app
COPY src /home/app/src
COPY pom.xml /home/app
RUN mvn -f /home/app/pom.xml clean package

# Production
FROM maven:3.8.1-jdk-11-slim AS production

WORKDIR /app

#RUN apk add --no-cache bash

COPY --from=builder /app/build/libs/*.jar app.jar

ENTRYPOINT ["java", "-jar", "app.jar"]
