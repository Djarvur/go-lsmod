package lsmod

import "fmt"

type modTainted uint32

// possible tainted flags of the module
const (
	TaintedNone modTainted = 0      // No tainted flag
	TaintedP    modTainted = 1      // (P): A module with a non-GPL license has been loaded, this includes modules with no license. Set by modutils >= 2.4.9 and module-init-tools.
	TaintedF    modTainted = 2      // (F): A module was force loaded by insmod -f. Set by modutils >= 2.4.9 and module-init-tools.
	TaintedS    modTainted = 4      // (S): Unsafe SMP processors: SMP with CPUs not designed for SMP.
	TaintedR    modTainted = 8      // (R): A module was forcibly unloaded from the system by rmmod -f.
	TaintedM    modTainted = 16     // (M): A hardware machine check error occurred on the system.
	TaintedB    modTainted = 32     // (B): A bad page was discovered on the system.
	TaintedU    modTainted = 64     // (U): The user has asked that the system be marked "tainted". This could be because they are running software that directly modifies the hardware, or for other reasons.
	TaintedD    modTainted = 128    // (D): The system has died.
	TaintedA    modTainted = 256    // (A): The ACPI DSDT has been overridden with one supplied by the user instead of using the one provided by the hardware.
	TaintedW    modTainted = 512    // (W): A kernel warning has occurred.
	TaintedC    modTainted = 1024   // (C): A module from drivers/staging was loaded.
	TaintedI    modTainted = 2048   // (I): The system is working around a severe firmware bug.
	TaintedO    modTainted = 4096   // (O): An out-of-tree module has been loaded.
	TaintedE    modTainted = 8192   // (E): An unsigned module has been loaded in a kernel supporting module signature.
	TaintedL    modTainted = 16384  // (L): A soft lockup has previously occurred on the system.
	TaintedK    modTainted = 32768  // (K): The kernel has been live patched.
	TaintedX    modTainted = 65536  // (X): Auxiliary taint, defined and used by for distros.
	TaintedT    modTainted = 131072 // (T): The kernel was built with the struct randomization plugin.
)

func parseTainted(t string) (modTainted, error) { // nolint: gocyclo
	switch t {
	case "(P)":
		return TaintedP, nil
	case "(F)":
		return TaintedF, nil
	case "(S)":
		return TaintedS, nil
	case "(R)":
		return TaintedR, nil
	case "(M)":
		return TaintedM, nil
	case "(B)":
		return TaintedB, nil
	case "(U)":
		return TaintedU, nil
	case "(D)":
		return TaintedD, nil
	case "(A)":
		return TaintedA, nil
	case "(W)":
		return TaintedW, nil
	case "(C)":
		return TaintedC, nil
	case "(I)":
		return TaintedI, nil
	case "(O)":
		return TaintedO, nil
	case "(E)":
		return TaintedE, nil
	case "(L)":
		return TaintedL, nil
	case "(K)":
		return TaintedK, nil
	case "(X)":
		return TaintedX, nil
	case "(T)":
		return TaintedT, nil
	}

	return 0, fmt.Errorf("unknown tainted flag %q", t)
}

func (t modTainted) String() string { // nolint: gocyclo
	switch t {
	case TaintedNone:
		return "()"
	case TaintedP:
		return "(P)"
	case TaintedF:
		return "(F)"
	case TaintedS:
		return "(S)"
	case TaintedR:
		return "(R)"
	case TaintedM:
		return "(M)"
	case TaintedB:
		return "(B)"
	case TaintedU:
		return "(U)"
	case TaintedD:
		return "(D)"
	case TaintedA:
		return "(A)"
	case TaintedW:
		return "(W)"
	case TaintedC:
		return "(C)"
	case TaintedI:
		return "(I)"
	case TaintedO:
		return "(O)"
	case TaintedE:
		return "(E)"
	case TaintedL:
		return "(L)"
	case TaintedK:
		return "(K)"
	case TaintedX:
		return "(X)"
	case TaintedT:
		return "(T)"
	}

	panic(fmt.Errorf("unknown tainted flag %d", t))
}
