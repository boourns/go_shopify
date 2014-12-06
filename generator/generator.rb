require 'erb'

class Overrides
  def self.overrides
    @overrides ||= JSON.parse(File.read("./data/overrides.json"))
  end

  def self.for(type, value)
    overrides["#{type}/#{value}"]
  end
end

class Generator
  class Property
    def initialize(prop, instance)
      @prop = prop
      @instance = instance
    end

    def name
      @prop["name"].gsub(' ', '')
    end

    def description
      @prop["description"]
    end

    def go_type
      example = @prop["examples"][0].values[0] if @prop["examples"][0]
      example = @instance || example
      return "string" if example.nil?

      return type_for(example)
    end

    def type_for(thing)
      return "string" if thing.nil?
      puts "thing is #{thing}, a #{thing.class}"
      t = Time.parse(thing) rescue nil
      if t
        return "time.Time"
      elsif thing.is_a?(Array) && thing == []
        return "[]interface{}"
      elsif thing.is_a?(Array)
        return "[]#{type_for(thing.first)}"
      elsif thing.is_a?(Integer)
        return "int64"
      elsif thing.is_a?(Float)
        return "float64"
      elsif thing.class.to_s =~ /\AShopifyAPI::(.*)\z/
        klass = Overrides.for("type", $1) || $1
        if !Generator.api_classes.include?(klass)
          return "interface{}"
          Generator.add_subclass(klass, thing)
        end
        return klass
      else
        return "string"
      end
    end

    def imports
      if go_type == "time.Time"
        ["time"]
      else
        []
      end
    end
  end

  class Action
    def initialize(klass, action)
      @klass = klass
      @name = action["name"]
    end

    def endpoint
      "/admin/#{@klass.name.underscore.pluralize}"
    end

    def render
      if File.exist? "./generator/methods/#{@name}.go.erb"
       template = File.read("./generator/methods/#{@name}.go.erb")
        
        ERB.new(template).result(binding)
      else
        puts "// TODO implement #{@klass.name}.#{@name}"
      end
    end

    def imports
      ["encoding/json", "fmt"]
    end
  end

  attr_reader :name

  def initialize(api, instance)
    @api = api
    @instance = instance
    @name = api["name"].gsub(' ', '')

    @properties = api["properties"].map do |p| 
      example = instance.send(p["name"].to_sym) rescue nil
      Property.new(p, example)
    end

    # properties have name:string, description:string, examples: [{parameter:value}]
    @actions = api["actions"].map { |a| Action.new(self, a) }
    # actions have name, heading, params: []
  end

  def render(template)
    filename = "./api/#{@name.underscore}.go"
    File.write(filename, ERB.new(template).result(binding))
  end

  def imports
    (@properties + @actions).map(&:imports).flatten.compact.sort.uniq
  end

  def self.api_classes=(names)
    @api_classes = names
  end

  def self.api_classes
    @api_classes
  end
end
