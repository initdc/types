task default: %w[fmt vet test]

task :fmt do
  sh "go fmt ./..."
end

task :vet do
  sh "go vet ./..."
end

task :test do
  sh "go test -cover -v ./..."
end

task :test_result do
  sh "go test -cover -v ./result"
end

task :test_option do
  sh "go test -cover -v ./option"
end
