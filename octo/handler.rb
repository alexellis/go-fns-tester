require 'octokit'


class Handler
	def run(req)
		user_name = ENV["Http_Path"].gsub(/#{Regexp.escape('/')}/, "")

		if user_name.length == 0 then
			return "To look-up the name of a GitHub user, pass their username as the HTTP path such as /alexellis"
		end

		client = Octokit::Client.new

		# Fetch a user
		user = client.user user_name
		return user.name
    end
end
