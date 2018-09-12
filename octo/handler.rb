require 'octokit'


class Handler
    def run(req)
		if ENV["Http_Path"].length == 0 then
			return "To look-up the name of a GitHub user, pass their username as the HTTP path such as /alexellis"
		end

		client = Octokit::Client.new

		# Fetch a user
		user = client.user ENV["Http_Path"]
		return user.name
    end
end
