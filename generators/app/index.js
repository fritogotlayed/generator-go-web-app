'use strict';
const Generator = require('yeoman-generator');
const chalk = require('chalk');
const yosay = require('yosay');

module.exports = class extends Generator {
  prompting() {
    // Have Yeoman greet the user.
    this.log(yosay(
      'Welcome to the terrific ' + chalk.red('generator-go-web-app') + ' generator!'
    ));

    // const prompts = [{
    //   type: 'confirm',
    //   name: 'someAnswer',
    //   message: 'Would you like to enable this option?',
    //   default: true
    // }];

    // return this.prompt(prompts).then(props => {
    //   // To access props later use this.props.someAnswer;
    //   this.props = props;
    // });
  }

  writing() {
    var files = [
      'config.go',
      'config.json',
      'glide.yaml',
      'handlers.go',
      'helpers.go',
      'logging.go',
      'main.go',
      'router.go',
      'routes.go',
      'server.go',
      'static/templates/index.html'
    ]
    for (var i = 0; i < files.length; i++) {
      var file = files[i];
      this.fs.copy(
        this.templatePath(file),
        this.destinationPath(file)
      );
    }
  }

  install() {
    // If there were npm / bower dependences uncomment this to run.
    // this.installDependencies();
  }
};
