# twister
The WIndow System for TERminals

Taking inspiration from [Twin](https://github.com/cosmos72/twin), a project that itself needs some more love, Twister aims to give you a psuedographical desktop environment in any supported terminal, providing multitasking, mouse support, simple management of subshells and virtual terminals, intelligent compositing of virtual terminals, multiple terminals, excellent mouse support, high customability and maximum performance in any host terminal, with a hopefully minimal resource footprint for the features provided. Twister is written in [Go](https://golang.org) and can take advantage of its powerful ability to run asynchronous processes, and uses zyedidia's fork of [tcell](https://github.com/zyedidia/tcell), a powerful terminal interface library for Go, which I also contribute to. If you're sure that the host terminal can support it, Twister will also provide true colour support, though it is not enabled by default.

Twister was written with Linux in mind, but should support BSD and MacOS straight out of the box. Due to the design, it should be possible to port it for Windows. Documentation is builtin and can easily be accessed. Likewise, configuration can be done from within a [dialog](https://linux.die.net/man/1/dialog)-like customization widget. TwistTerm, Twister's internal terminal emulator system should naturally and transparently support utf-8 input and output: when Twister detects that its host terminal does not support utf-8 output it will try to gloss it to a supported character and if this fails it will print one character per unsupported character ( unlike some programs that print multiple characters for unsupported wide characters ), otherwise it will input and output in utf-8.

Twister should have absolutely no problems with running inside screen or tmux, and should be able to flawlessly run screen and tmux in its terminal emulator. Like screen and tmux. It should theoretically be able to run in twin, although this is more of a hope than a design goal. While not currently recommended, Twister can also run in the Linux console (tty devices), though it's subject to the limitations thereof. ( However, Twister is planned to eventually use hot-switching of the Linux console fonts as a workaround to the 256 character limitation of the Linux console, though this feature will never be flawless. ) If you want to use mouse support in the linux console, you will need to install [gpm](http://www.nico.schottelius.org/software/gpm/).

Screenshots to come when development has progressed enough to have them.
