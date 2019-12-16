module Dsl
  def video(name)
    
  end

  def uniform(name, value)
    @uniforms ||= {}
    @uniforms[name] = value
  end
end

extend Dsl

uniform :banana, 'rama'
puts @uniforms
