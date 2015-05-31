require 'bundler'
require 'rubygems'
require 'active_record'

RSpec.configure do |config|
    config.filter_run :focus
    config.run_all_when_everything_filtered = true
    config.default_formatter = 'doc'
    config.order = :random

    Kernel.srand config.seed
end

ActiveRecord::Base.establish_connection(
    :adapter => 'mysql',
    :host => 'localhost',
    :database => 'converter'
)

class Members < ActiveRecord::Base
end