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

DemoInfo = Struct.new(:mapName, :serverName, :durationInSec, :durationInTick, :durationInFrame)
Team = Struct.new(:id, :name, :players)
Player = Struct.new(:steamID, :gameID, :name)

get '/' do
  haml :index, format: :html5
end

post '/upload' do
  tempfile = params[:file][:tempfile]
  stdout, _stderr, _status = Open3.capture3('../out/bin/cs2-reviewer', tempfile.path)
  json = Oj.load(stdout)

  @demoInfo = DemoInfo.new(**json['demoInfo'])
  @teams = []
  @teams[0] = Team.new(json['teams'][0]['id'], json['teams'][0]['name'])
  @teams[1] = Team.new(json['teams'][1]['id'], json['teams'][1]['name'])

  @teams[0].players = json['teams'][0]['players'].map do |player|
    Player.new(**player)
  end
  @teams[1].players = json['teams'][1]['players'].map do |player|
    Player.new(**player)
  end

  @rounds = json['rounds']

  haml :upload, format: :html5
end
