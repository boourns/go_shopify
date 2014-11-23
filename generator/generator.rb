require 'erb'

class Generator
  class Property
    def initialize(prop)
      @prop = prop
    end

    def name
      @prop["name"].gsub(' ', '')
    end

    def description
      @prop["description"]
    end

    def go_type
      return "string" if @prop["examples"][0].nil?

      example = @prop["examples"][0].values[0]
      t = Time.parse(example) rescue nil
      if t
        return "time.Time"
      elsif example.is_a?(Integer)
        return "int64"
      elsif example.is_a?(Float)
        return "float64"
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
      ["/admin/#{@klass.name.underscore.pluralize}.json"].join("")
    end

    def method
      case @name
      when 'index', 'show'
        'GET'
      when 'create', 'destroy'
        'POST'
      when 'update'
        'PATCH'
      else
        raise "Don't know HTTP method for #{@name}"
      end
    end

    def render
      if File.exist? "./generator/methods/#{@name}.go.erb"
       template = File.read("./generator/methods/#{@name}.go.erb")
        
        ERB.new(template).result(binding)
      else
        "// TODO implement #{@klass.name}.#{@name}"
      end
    end

    def imports
      ["encoding/json", "fmt"]
    end
  end

  attr_reader :name

  def initialize(api)
    @api = api
    @name = api["name"].gsub(' ', '')
    @properties = api["properties"].map { |p| Property.new(p) }

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
end
