#
# Requiring Your Gemfile
#

require 'rubygems'
require 'bundler/setup'

Bundler.require(:default)

#
# Require all the module you're creating here
#

loader = Zeitwerk::Loader.new
loader.push_dir('./lib')
loader.setup

#
# v  Write your main loop downhere v
#

require 'open3'

get '/' do
  haml :index, format: :html5
end

post '/upload' do
  tempfile = params[:file][:tempfile]
  stdout, _stderr, _status = Open3.capture3('../out/bin/cs2-reviewer', tempfile.path)
  stdout
end
