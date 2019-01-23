package lsmod

import "fmt"

type modTained uint32

const (
	TainedNone modTained = 0      // No tained flag
	TainedP    modTained = 1      // (P): A module with a non-GPL license has been loaded, this includes modules with no license. Set by modutils >= 2.4.9 and module-init-tools.
	TainedF    modTained = 2      // (F): A module was force loaded by insmod -f. Set by modutils >= 2.4.9 and module-init-tools.
	TainedS    modTained = 4      // (S): Unsafe SMP processors: SMP with CPUs not designed for SMP.
	TainedR    modTained = 8      // (R): A module was forcibly unloaded from the system by rmmod -f.
	TainedM    modTained = 16     // (M): A hardware machine check error occurred on the system.
	TainedB    modTained = 32     // (B): A bad page was discovered on the system.
	TainedU    modTained = 64     // (U): The user has asked that the system be marked "tainted". This could be because they are running software that directly modifies the hardware, or for other reasons.
	TainedD    modTained = 128    // (D): The system has died.
	TainedA    modTained = 256    // (A): The ACPI DSDT has been overridden with one supplied by the user instead of using the one provided by the hardware.
	TainedW    modTained = 512    // (W): A kernel warning has occurred.
	TainedC    modTained = 1024   // (C): A module from drivers/staging was loaded.
	TainedI    modTained = 2048   // (I): The system is working around a severe firmware bug.
	TainedO    modTained = 4096   // (O): An out-of-tree module has been loaded.
	TainedE    modTained = 8192   // (E): An unsigned module has been loaded in a kernel supporting module signature.
	TainedL    modTained = 16384  // (L): A soft lockup has previously occurred on the system.
	TainedK    modTained = 32768  // (K): The kernel has been live patched.
	TainedX    modTained = 65536  // (X): Auxiliary taint, defined and used by for distros.
	TainedT    modTained = 131072 // (T): The kernel was built with the struct randomization plugin.

)

func parseTained(t string) (modTained, error) { // nolint: gocyclo
	switch t {
	case "(P)":
		return TainedP, nil
	case "(F)":
		return TainedF, nil
	case "(S)":
		return TainedS, nil
	case "(R)":
		return TainedR, nil
	case "(M)":
		return TainedM, nil
	case "(B)":
		return TainedB, nil
	case "(U)":
		return TainedU, nil
	case "(D)":
		return TainedD, nil
	case "(A)":
		return TainedA, nil
	case "(W)":
		return TainedW, nil
	case "(C)":
		return TainedC, nil
	case "(I)":
		return TainedI, nil
	case "(O)":
		return TainedO, nil
	case "(E)":
		return TainedE, nil
	case "(L)":
		return TainedL, nil
	case "(K)":
		return TainedK, nil
	case "(X)":
		return TainedX, nil
	case "(T)":
		return TainedT, nil
	}
	return 0, fmt.Errorf("unknown tained flag %q", t)
}

func (t modTained) String() string { // nolint: gocyclo
	switch t {
	case TainedNone:
		return "()"
	case TainedP:
		return "(P)"
	case TainedF:
		return "(F)"
	case TainedS:
		return "(S)"
	case TainedR:
		return "(R)"
	case TainedM:
		return "(M)"
	case TainedB:
		return "(B)"
	case TainedU:
		return "(U)"
	case TainedD:
		return "(D)"
	case TainedA:
		return "(A)"
	case TainedW:
		return "(W)"
	case TainedC:
		return "(C)"
	case TainedI:
		return "(I)"
	case TainedO:
		return "(O)"
	case TainedE:
		return "(E)"
	case TainedL:
		return "(L)"
	case TainedK:
		return "(K)"
	case TainedX:
		return "(X)"
	case TainedT:
		return "(T)"
	}
	panic(fmt.Errorf("unknoiwn tained flag %d", t))
}
