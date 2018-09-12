require 'octokit'


class Handler
    def run(req)
        
	client = Octokit::Client.new

	# Fetch a user
	user = client.user ENV["Http_Path"]
	return user.name
    end
end
