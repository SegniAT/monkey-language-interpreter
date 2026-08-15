/**
 * @file Monkey grammar for tree-sitter
 * @author SegniAT <se.segni.adeba@gmail.com>
 * @license MIT
 */

/// <reference types="tree-sitter-cli/dsl" />
// @ts-check

export default grammar({
  name: "monkey",

  supertypes: $ => [
    $.expression,
  ],

  word: $ => $.identifier,

  rules: {
    source_file: $ => repeat($._statement),

    _statement: $ => choice(
      $.let_statement,
      $.return_statement,
      $.expression_statement
    ),

    let_statement: $ => seq(
      'let',
      field('name', $.identifier),
      '=',
      field('value', $.expression),
      optional(';')
    ),

    return_statement: $ => seq(
      'return',
      field('value', $.expression),
      optional(';')
    ),

    expression_statement: $ => seq(
      $.expression,
      optional(';')
    ),

    expression: $ => choice(
      $.identifier,
      $.integer,
      $.string,
      $.boolean,

      $.unary_expression,
      $.binary_expression,
      $.paren_expression,

      $.call_expression,
      $.index_expression,

      $.if_expression,
      $.function_literal,
      $.array_literal,
      $.hash_literal,
    ),

    identifier: _ => /[a-zA-Z_]+/,
    integer: _ => /\d+/,
    string: _ => token(seq('"', /[^"]*/, '"')), // no escapes
    boolean: _ => choice("true", "false"),

    unary_expression: $ => prec(5, choice(
      seq('-', $.expression),
      seq('!', $.expression),
    )),

    binary_expression: $ => choice(
      prec.left(1, seq(field('left', $.expression), '==', field('right', $.expression))),
      prec.left(1, seq(field('left', $.expression), '!=', field('right', $.expression))),
      prec.left(2, seq(field('left', $.expression), '<', field('right', $.expression))),
      prec.left(2, seq(field('left', $.expression), '>', field('right', $.expression))),
      prec.left(3, seq(field('left', $.expression), '+', field('right', $.expression))),
      prec.left(3, seq(field('left', $.expression), '-', field('right', $.expression))),
      prec.left(4, seq(field('left', $.expression), '*', field('right', $.expression))),
      prec.left(4, seq(field('left', $.expression), '/', field('right', $.expression))),
    ),

    paren_expression: $ => seq('(', $.expression, ')'),

    call_expression: $ => prec.left(7, seq(
      field('function', $.expression),
      field('arguments', $.argument_list)
    )),

    index_expression: $ => prec.left(8, seq(
      field('object', $.expression),
      '[',
      field('index', $.expression),
      ']',
    )),

    argument_list: $ => seq('(', optional($._expression_list), ')'),

    _expression_list: $ => seq($.expression, repeat(seq(',', $.expression))),

    block: $ => seq('{', repeat($._statement), '}'),

    if_expression: $ => seq(
      'if',
      '(',
      field('condition', $.expression),
      ')',
      field('consequence', $.block),
      optional(seq('else', field('alternative', $.block))),
    ),

    function_literal: $ => seq(
      'fn',
      field('parameters', $.parameter_list),
      field('body', $.block)
    ),

    parameter_list: $ => seq(
      '(',
      optional(seq(
        field('name', $.identifier),
        repeat(seq(',', field('name', $.identifier)))
      )),
      ')'
    ),

    array_literal: $ => seq('[', optional($._expression_list), ']'),

    hash_literal: $ => seq(
      '{',
      optional(seq($.hash_pair, repeat(seq(',', $.hash_pair)), optional(','))),
      '}',
    ),

    hash_pair: $ => seq(
      field('key', $.expression),
      ':',
      field('value', $.expression),
    )
  }
});
