require 'rake'
require 'erb'
require 'json'
require 'active_support/all'
require 'pp'
require './generator/generator'

task :default => :generate

task :generate do |t|
  schema = JSON::parse(File.read("./data/api_schema.json"))

  schema.each do |klass|
    Generator.new(klass).render(File.read("./generator/template.go.erb"))
  end

  `go fmt`
end

