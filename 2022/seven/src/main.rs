use std::fs;
use std::str::FromStr;

fn main() {
    let contents = fs::read_to_string("input.txt").expect("Something went wrong reading the file");
    let lines = contents.lines().collect::<Vec<&str>>();

    for (_, c) in lines.iter().enumerate() {
        let split = c.split(" ").collect::<Vec<&str>>();
        let is_command = split[0] == "$";
        if is_command {
            let command = Command::from_str(split[1]);
            match command {
                Ok(Command::Ls(s)) => println!("ls {}", s),
                Ok(Command::Cd(_)) => handle_cd(split[2].to_string()), 
                Err(_) => println!("Invalid command"),
            }
        } else {
            println!("File: {}", c);
        }
    }
}

fn handle_cd(target: String) {
    let target = CdTo::from_str(&target);
    match target {
        Ok(CdTo::Root) => println!("cd /"),
        Ok(CdTo::Parent) => println!("cd .."),
        Ok(CdTo::Child(s)) => println!("cd {}", s),
        Err(_) => println!("Invalid cd target"),
    }
}

enum Command {
    Ls(String),
    Cd(String),
}

enum CdTo {
    Root,
    Parent,
    Child(String),
}

impl FromStr for CdTo {
    type Err = ();

    fn from_str(s: &str) -> Result<Self, Self::Err> {
        match s {
            "/" => Ok(CdTo::Root),
            ".." => Ok(CdTo::Parent),
            _ => Ok(CdTo::Child(String::from(s))),
        }
    }
}

impl FromStr for Command {
    type Err = ();

    fn from_str(s: &str) -> Result<Self, Self::Err> {
        match s {
            "ls" => Ok(Command::Ls(String::from(s))),
            "cd" => Ok(Command::Cd(String::from(s))),
            _ => Err(()),
        }
    }
}

struct Node {
    children: Vec<Node>,
    parent: Option<Box<Node>>,
    node_type: NodeType,
    data: String,
}

enum NodeType {
    Directory,
    File,
}

impl Node {
    fn new(node_type: NodeType, data: String) -> Node {
        Node {
            children: Vec::new(),
            parent: None,
            node_type,
            data,
        }
    }

    fn add_child(&mut self, child: Node) {
        self.children.push(child);
    }

    fn add_parent(&mut self, parent: Node) {
        self.parent = Some(Box::new(parent));
    }
}
