3.2 - Strategies of Evaluation (Pg.104) --

The notion of an interpreter as something that doesn’t leave executable
artifacts behind (in contrast to a compiler, which does just that) gets fuzzy real fast when looking
at the implementations of real-world and highly-optimized programming languages.

With that said, the most obvious and classical choice of what to do with the AST is to just
interpret it. Traverse the AST, visit each node and do what the node signifies: print a string, add
two numbers, execute a function’s body - all on the fly. Interpreters working this way are called
“tree-walking interpreters” and are the archetype of interpreters. Sometimes their evaluation
step is preceded by small optimizations that rewrite the AST (e.g. remove unused variable
bindings) or convert it into another intermediate representation (IR) that’s more suitable for
recursive and repeated evaluation.

Other interpreters also traverse the AST, but instead of interpreting the AST itself they first
convert it to bytecode. Bytecode is another IR of the AST and a really dense one at that. The
exact format and of which opcodes (the instructions that make up the bytecode) it’s composed
of varies and depends on the guest and host programming languages. In general though, the
opcodes are pretty similar to the mnemonics of most assembly languages; it’s a safe bet to
say that most bytecode definitions contain opcodes for push and pop to do stack operations.
But bytecode is not native machine code, nor is it assembly language. It can’t and won’t be
executed by the operating system and the CPU of the machine the interpreter is running on.
Instead it’s interpreted by a virtual machine, that’s part of the interpreter. Just like VMWare
and VirtualBox emulate real machines and CPUs, these virtual machines emulate a machine
that understands this particular bytecode format. This approach can yield great performance
benefits.

A variation of this strategy doesn’t involve an AST at all. Instead of building an AST the parser
emits bytecode directly. Now, are we still talking about interpreters or compilers? Isn’t emit-
ting bytecode that gets then interpreted (or should we say “executed”?) a form of compilation?
I told you: the line becomes blurry. And to make it even more fuzzy, consider this: some im-
plementations of programming languages parse the source code, build an AST and convert this
AST to bytecode. But instead of executing the operations specified by the bytecode directly in
a virtual machine, the virtual machine then compiles the bytecode to native machine code, right
before its executed - just in time. That’s called a JIT (for “just in time”) interpreter/compiler.

Others skip the compilation to bytecode. They recursively traverse the AST but before executing
a particular branch of it the node is compiled to native machine code. And then executed. Again,
(Pg.104)
“just in time”.

A slight variation of this is a mixed mode of interpretation where the interpreter recursively
evaluates the AST and only after evaluating a particular branch of the AST multiple times does
it compile the branch to machine code.

** The choice of which strategy to choose largely depends on performance and portability needs,
the programming language that’s being interpreted and how far you’re willing to go. A tree-
walking interpreter that recursively evaluates an AST is probably the slowest of all approaches,
but easy to build, extend, reason about and as portable as the language it’s implemented in.

** An interpreter that compiles to bytecode and uses a virtual machine to evaluate said bytecode is
going to be a lot faster. But more complicated and harder to build, too. Throw JIT compilation
to machine code into the mix and now you also need to support multiple machine architectures
if you want the interpreter to work on both ARM and x86 CPUs.

[Different real-world implementations]:

All of these approaches can be found in real-world programming languages. And most of the time
the chosen approach changed with the lifetime of the language. Ruby is a great example here.
Up to and including version 1.8 the interpreter was a tree-walking interpreter, executing the
AST while traversing it. But with version 1.9 came the switch to a virtual machine architecture.
Now the Ruby interpreter parses source code, builds an AST and then compiles this AST into
bytecode, which gets then executed in a virtual machine. The increase in performance was huge.

The WebKit JavaScript engine JavaScriptCore and its interpreter named “Squirrelfish” also
used AST walking and direct execution as its approach. Then in 2008 came the switch to a
virtual machine and bytecode interpretation. Nowadays the engine has four (!) different stages
of JIT compilation, which kick in at different times in the lifetime of the interpreted program –
depending on which part of the program needs the best performance.

Another example is Lua. The main implementation of the Lua programming language started
out as an interpreter that compiles to bytecode and executes the bytecode in a register-based
virtual machine. 12 years after its first release another implementation of the language was
born: LuaJIT. The clear goal of Mike Pall, the creator of LuaJIT, was to create the fastest
Lua implementation possible. And he did. By JIT compiling a dense bytecode format to
highly-optimized machine code for different architectures the LuaJIT implementation beats the
original Lua in every benchmark. And not just by a tiny bit, no; it’s sometimes 50 times faster.

3.4 - Representing Objects (Pg.106) --

REFERENCE: The point is this: there are a lot of different ways to represent values of the interpreted lan-
guages in the host language. The best (and maybe the only) way to learn about these different
representations is to actually read through the source code of some popular interpreters. I
heartily recommended the Wren source code, which includes two types of value representation,
enabled/disabled by using a compiler flag.

Besides the representation of values inside the host language there is also the matter of how to
expose these values and their representation to the user of the interpreted language. What does
the “public API” of these values look like?

Java, for example, offers both “primitive data types” (int, byte, short, long, float, double,
boolean, char) and reference types to the user. The primitive data types do not have a huge
representation inside the Java implementation, they closely map to their native counterparts.
Reference types on the other hand are references to compound data structures defined in the
host language.

(Pg.107)

In Ruby the user doesn’t have access to “primitive data types”, nothing like a native value
type exists because everything is an object and thus wrapped inside an internal representation.
Internally Ruby doesn’t distinguish between a byte and an instance of the class Pizza: both are
the same value type, wrapping different values.
