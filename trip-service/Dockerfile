FROM maven:3.9.16-eclipse-temurin-21 AS builder

WORKDIR /app

COPY trip-service/pom.xml .
RUN mvn dependency:go-offline

COPY trip-service/src ./src
COPY proto ./proto

RUN mvn clean package -DskipTests

FROM eclipse-temurin:21-jre-alpine

WORKDIR /app

COPY --from=builder /app/target/*.jar app.jar

EXPOSE 9090 8080

CMD ["java", "-jar", "app.jar"]