require 'rake'
require 'erb'
require 'json'
require 'active_support/all'
require 'pp'
require './generator/generator'
require 'shopify_api'
require 'byebug'

task :default => :generate

task :generate do |t|

  if !ENV["SHOPIFY_API_TOKEN"] || !ENV["SHOPIFY_API_SECRET"] || !ENV["SHOPIFY_API_HOST"]
    raise "Set SHOPIFY_API_TOKEN SHOPIFY_API_SECRET SHOPIFY_API_HOST"
  end

  target = URI.parse(ENV["SHOPIFY_API_HOST"])
  target.user = ENV["SHOPIFY_API_TOKEN"]
  target.password = ENV["SHOPIFY_API_SECRET"]
  target.path = "/admin"

  ShopifyAPI::Base.site = target.to_s

  schema = JSON::parse(File.read("./data/api_schema.json"))

  Generator.api_classes = schema.map { |klass| klass["name"] }.sort

  schema.each do |klass|
    name = klass["name"]
    obj = ShopifyAPI::const_get(name).first() rescue nil
    Generator.new(klass, obj).render(File.read("./generator/template.go.erb"))
  end

  `go fmt`
end

