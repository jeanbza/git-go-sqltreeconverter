require File.join(File.dirname(__FILE__), '../spec_helper')

describe 'the users api', :type => :feature do
    describe 'GET /users' do
        it 'returns a list of users ready for the tree javascript' do
            # resp = HTTParty.get('http://localhost:8080/')
            #
            # expect(resp.code).to eq 200
            # expect(JSON.parse(resp.body)).to eq [{"name" => 'Super Glue'}, {"name" => 'Kool-Aide'}]
        end
    end
end