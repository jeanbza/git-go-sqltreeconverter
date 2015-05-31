require_relative 'spec_helper'

describe 'sql converter', :type => :feature do
    describe 'multiple roots' do
        it 'works' do
            root_file = 'input_multiple_roots_complex'
            setup_database(root_file)

            compile_converter
            run_converter

            members = Members.all

            members.each do |member|
                puts "#{member.id} #{member.parent_id}"
            end

            expect(4).to eq 4
        end
    end
end

def test_dir
    File.dirname(__FILE__)
end

def app_dir
    "#{test_dir}/../.."
end

def setup_database(root_file)
    puts `mysql -uroot --database converter -e 'source #{test_dir}/fixtures/#{root_file}.sql'`
end

def compile_converter
    puts `cd #{app_dir} && go build 2>&1`
end

def run_converter(input_file_name = 'input_multiple_roots_complex')
    puts `#{app_dir}/git-go-sqltreeconverter --input #{test_dir}/fixtures/#{input_file_name}.sql --output /tmp/output.sql --regex "(\\w)\t(\\w)" --target converter.members`
end