package main

import (
	"blackjack/internal/deck"
	"errors"
	"flag"
	"fmt"
	"os"
)

func main() {
	os.Exit(run(os.Args[1:]))
}

func run(args []string) int {
	if len(args) == 0 {
		printUsage()
		return 1
	}

	switch args[0] {
	case "-h", "--help", "help":
		printUsage()
		return 0
	case "new":
		return runNew(args[1:])
	case "shuffle":
		return runShuffle(args[1:])
	case "deal":
		return runDeal(args[1:])
	default:
		fmt.Fprintf(os.Stderr, "Error: unknown subcommand %q\n", args[0])
		printUsage()
		return 1
	}
}

func runNew(args []string) int {
	fs := flag.NewFlagSet("new", flag.ContinueOnError)
	fs.SetOutput(os.Stdout)
	outPath := fs.String("out", "assets/my_cards.txt", "output file path")
	fs.Usage = func() {
		fmt.Fprintf(fs.Output(), "Usage: deck new [--out <path>]\n")
		fs.PrintDefaults()
	}

	if err := fs.Parse(args); err != nil {
		if errors.Is(err, flag.ErrHelp) {
			return 0
		}

		return 1
	}

	if fs.NArg() > 0 {
		fmt.Fprintln(os.Stderr, "Error: unexpected positional arguments for new")
		fs.Usage()
		return 1
	}

	if err := deck.NewToFile(*outPath); err != nil {
		fmt.Fprintf(os.Stderr, "Error: could not create deck file: %v\n", err)
		return 1
	}

	fmt.Printf("Deck saved to %s\n", *outPath)
	return 0
}

func runShuffle(args []string) int {
	fs := flag.NewFlagSet("shuffle", flag.ContinueOnError)
	fs.SetOutput(os.Stdout)
	inPath := fs.String("in", "", "input file path")
	outPath := fs.String("out", "", "output file path (defaults to --in)")
	fs.Usage = func() {
		fmt.Fprintf(fs.Output(), "Usage: deck shuffle --in <path> [--out <path>]\n")
		fs.PrintDefaults()
	}

	if err := fs.Parse(args); err != nil {
		if errors.Is(err, flag.ErrHelp) {
			return 0
		}

		return 1
	}

	if fs.NArg() > 0 {
		fmt.Fprintln(os.Stderr, "Error: unexpected positional arguments for shuffle")
		fs.Usage()
		return 1
	}

	if *inPath == "" {
		fmt.Fprintln(os.Stderr, "Error: --in is required")
		fs.Usage()
		return 1
	}

	targetPath := *outPath
	if targetPath == "" {
		targetPath = *inPath
	}

	if err := deck.ShuffleFile(*inPath, targetPath); err != nil {
		fmt.Fprintf(os.Stderr, "Error: could not shuffle deck: %v\n", err)
		return 1
	}

	fmt.Printf("Shuffled deck saved to %s\n", targetPath)
	return 0
}

func runDeal(args []string) int {
	fs := flag.NewFlagSet("deal", flag.ContinueOnError)
	fs.SetOutput(os.Stdout)
	inPath := fs.String("in", "", "input file path")
	handSize := fs.Int("hand", 5, "hand size")
	fs.Usage = func() {
		fmt.Fprintf(fs.Output(), "Usage: deck deal --in <path> [--hand <n>]\n")
		fs.PrintDefaults()
	}

	if err := fs.Parse(args); err != nil {
		if errors.Is(err, flag.ErrHelp) {
			return 0
		}

		return 1
	}

	if fs.NArg() > 0 {
		fmt.Fprintln(os.Stderr, "Error: unexpected positional arguments for deal")
		fs.Usage()
		return 1
	}

	if *inPath == "" {
		fmt.Fprintln(os.Stderr, "Error: --in is required")
		fs.Usage()
		return 1
	}

	hand, remaining, err := deck.DealFile(*inPath, *handSize)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		return 1
	}

	fmt.Printf("Hand (%d cards):\n", len(hand))
	for i, card := range hand {
		fmt.Printf("%d %s\n", i, card)
	}
	fmt.Printf("Remaining deck size: %d\n", len(remaining))
	return 0
}

func printUsage() {
	fmt.Println("Usage:")
	fmt.Println("  deck new [--out <path>]")
	fmt.Println("  deck shuffle --in <path> [--out <path>]")
	fmt.Println("  deck deal --in <path> [--hand <n>]")
	fmt.Println()
	fmt.Println("Use \"deck <subcommand> --help\" for subcommand options.")
}
