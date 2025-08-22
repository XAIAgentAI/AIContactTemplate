// SPDX-License-Identifier: MIT
pragma solidity ^0.8.22;

enum StakingType {
    ShortTerm,
    LongTerm,
    Free
}

enum MachineStatus {
    Normal,
    Blocking
}

enum NotifyType {
    ContractRegister,
    MachineRegister,
    MachineUnregister,
    MachineOnline,
    MachineOffline
}
