require 'octokit'


class Handler
    def run(req)
	client = Octokit::Client.new

	# Fetch a user
	user = client.user req
	return user.name
    end
end
