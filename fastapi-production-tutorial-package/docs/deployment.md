# Deployment

## Local
```bash
docker compose up --build
curl http://localhost:8000/health
```

## Jenkins
```text
Git -> Jenkins -> lint -> test -> Docker build -> ECR -> EKS rollout
```
Configure AWS/Jenkins credentials through a credential store or IAM role. Never commit access keys.

## EKS
```bash
aws eks update-kubeconfig --region ap-south-1 --name <cluster>
kubectl apply -f k8s/
kubectl rollout status deployment/fastapi-api -n fastapi
```

## Production edge
```text
Internet -> Route53 -> CloudFront/WAF -> ALB -> EKS Ingress -> FastAPI
```
