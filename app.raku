use Cro::HTTP::Router;
use Cro::HTTP::Server;
use Cro::WebApp::Template;

my $application = route {


  post -> 'hook', {
    request-body -> %json {
      say "hit";
      content 'application/json', %json; 
    }
  }

}

my Cro::Service $service = Cro::HTTP::Server.new:
    :host(%*ENV<DSCI_HOST> || "127.0.0.1"), :port(%*ENV<DSCI_PORT> || 3333), :$application;

$service.start;

react whenever signal(SIGINT) {
    $service.stop;
    exit;
}

