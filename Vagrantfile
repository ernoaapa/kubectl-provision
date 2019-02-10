# -*- mode: ruby -*-
# vi: set ft=ruby :

Vagrant.configure("2") do |config|
  (1..2).each do |i|
    name = "node-#{i}"
    config.vm.define name do |node|
      config.vm.box = "ubuntu/bionic64"

      # Use the VM unique id as hostname to make it unique
      hostname_file = File.join(".vagrant", name)
      File.open(hostname_file, "w") {|f| f.write("vagrant-#{SecureRandom.uuid().tr("-","")}")} unless File.file?(hostname_file)
      hostname = File.read(hostname_file)
      config.vm.hostname = hostname

      config.trigger.before :destroy do |trigger|
        trigger.warn = "Delete node from the cluster"
        trigger.run = {inline: "kubectl delete node #{hostname}"}
        trigger.exit_codes = [0, 1] # Ignore failures
      end
    end
  end
end
