
//! Module lexer is generated by GoGLL. Do not edit.

use crate::token;

use std::{fs, io};
use std::rc::Rc;

type State = usize;

const NULL_STATE: State = usize::MAX ;

/**
Lexer contains both the input Vec<char> and the Vec<token::Token>
parsed from the input
*/
pub struct Lexer {
	/// i is the input vector of char
	i: Rc<Vec<char>>,

	/// tokens is the vector of tokens constructed by the lexer from I
	pub tokens: Vec<Rc<token::Token>>
}

impl Lexer {
	/**
	new_file constructs a Lexer created from the input file, fname. 

	If the input file is a markdown file new_file process treats all text outside
	code blocks as whitespace. All text inside code blocks are treated as input text.

	If the input file is a normal text file new_file treats all text in the inputfile
	as input text.
	*/
	#[allow(dead_code)]
	pub fn new_file(fname: &String) -> io::Result<Rc<Lexer>> {
		let i = Rc::new(load_file(fname)?);
		Ok(Lexer::new(i))
	}

	/**
	new constructs a Lexer from a Vec<char>. 
	
	All contents of the input are treated as input text.
	*/
	pub fn new(input: Rc<Vec<char>>) -> Rc<Lexer> {
		let mut lex = Lexer{
			i:      input.clone(),
			tokens: Vec::new(),
		};
		let mut lext = 0;
		while lext < lex.i.len() {
			while lext < lex.i.len() && lex.i[lext].is_whitespace() {
				lext += 1
			}
			if lext < lex.i.len() {
				let tok = lex.scan(lext);
				lext = tok.rext;
				if !tok.suppress() {
					lex.add_token(tok)
				}
			}
		}
		lex.add(token::Type::EOF, input.len(), input.len());
		Rc::new(lex)
	}

	fn add(&mut self, t: token::Type, lext: usize, rext: usize) {
		self.add_token(token::new(t, lext, rext, &self.i))
	}
	
	fn add_token(&mut self, tok: Rc<token::Token>) {
		self.tokens.push(tok)
	}
	
	fn scan(&mut self, i: usize) -> Rc<token::Token> {
		let mut s: State = 0;
		let mut typ = token::Type::Error;
		let mut rext = i;

		while s != NULL_STATE {
			if rext >= self.i.len() {
				typ = ACCEPT[s];
				s = NULL_STATE
			} else {
				typ = ACCEPT[s];
				s = NEXT_STATE[s](self.i[rext]);
				if s != NULL_STATE || typ == token::Type::Error {
					rext += 1
				}
			}
		}
		return token::new(typ, i, rext, &self.i)
	}

	/// get_line_column returns the (line, column) of char[i] in the input
	#[allow(dead_code)]
	pub fn get_line_column(&self, i: usize) -> (usize, usize) {
		let mut line = 1;
		let mut col = 1;
		let mut j = 0;
		while j < i {
			match self.i[j] {
			'\n' => {
				line += 1;
				col = 1
			},
			'\t' => col += 4,
			_ => col += 1
			}
			j += 1
		}
		(line, col)
	}
	
	/// get_line_column_of_token returns the (line, column) of token[i] 
	/// in the input
	#[allow(dead_code)]
	pub fn get_line_column_of_token(&self, i: usize) -> (usize, usize) {
		self.get_line_column(self.tokens[i].lext)
	}

	// get_string returns the input string from the left extent of Token[lext] to
	// the right extent of Token[rext]
	#[allow(dead_code)]
	pub fn get_string(&self, lext: usize, rext: usize) -> String {
		let lext = self.tokens[lext].lext;
		let rext = self.tokens[rext].rext;
		self.i[lext..rext].iter().collect::<String>()
	}
	
	}
/*** End of Lexer implementation ***/


fn load_file(fname: &String) -> io::Result<Vec<char>> {
	let input = fs::read_to_string(fname)?;
	let mut input: Vec<char> = input.chars().collect();
	if fname.ends_with(".md") {
        load_md(&mut input)?;
        Ok(input)
	} else {
		Ok(input)
	}
}

fn load_md(input: &mut Vec<char>) -> io::Result<()> {
    let mut i = 0;
    let mut text = true;
    while i < input.len() {
        if i <= input.len() - 3 && 
        || -> bool { input[i] == '`' && input[i+1] == '`' && input[i+2] == '`' }()
        {
            text = !text;
            for j in i..i+3 {
                input[j] = ' ';
            }
            i += 3;
        }
        if i < input.len() {
            if text {
                match input[i] {
                    '\n' => input[i] = '\n',
                    _ => input[i] = ' ',
                }
            }
            i += 1;
        }
    }
    Ok(())
}

#[allow(dead_code)]
fn any(r: char, set: &'static [char]) -> bool {
	for r1 in set.iter() {
		if &r == r1 {
			return true
		}
	}
	return false
}

#[allow(dead_code)]
fn not(r: char, set: &'static [char]) -> bool {
	for r1 in set.iter() {
		if &r == r1 {
			return false
		}
	}
	return true
}

static ACCEPT: [token::Type; 8] = [ 
    token::Type::Error, 
    token::Type::Error, 
    token::Type::T_2, 
    token::Type::Error, 
    token::Type::Error, 
    token::Type::Error, 
    token::Type::T_1, 
    token::Type::T_0, 
];

pub type NextFun = dyn Fn(char) -> State + Sync;

static NEXT_STATE: &'static [&NextFun; 8] = &[  
	// Set0 
	&|c| -> State {  
        if c == '/' { return 1 }; 
        if c.is_alphabetic() { return 2 }; 
        NULL_STATE
	}, 
	// Set1 
	&|c| -> State {  
        if c == '*' { return 3 }; 
        if c == '/' { return 4 }; 
        NULL_STATE
	}, 
	// Set2 
	&|c| -> State {  
        if c.is_alphabetic() { return 2 }; 
        if c.is_numeric() { return 2 }; 
        NULL_STATE
	}, 
	// Set3 
	&|c| -> State {  
        if c == '*' { return 5 }; 
        if not(c, &['*']) { return 3 }; 
        NULL_STATE
	}, 
	// Set4 
	&|c| -> State {  
        if c == '\n' { return 6 }; 
        if not(c, &['\n']) { return 4 }; 
        NULL_STATE
	}, 
	// Set5 
	&|c| -> State {  
        if c == '/' { return 7 }; 
        NULL_STATE
	}, 
	// Set6 
	&|_| -> State {  
        NULL_STATE
	}, 
	// Set7 
	&|_| -> State {  
        NULL_STATE
	}, 
];
