// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract Attendance {

    // Struct to hold attendance records
    struct AttendanceRecord {
        uint256 id;
        uint256 employerId;
        uint256 checkInTime;
        string notes;
    }

    // State variables
    address public owner;
    mapping(address => bool) private authorizedEntities;

    AttendanceRecord[] private attendanceRecords;

    // Events
    event AttendanceRecorded(uint256 indexed id, uint256 indexed employeeId, uint256 checkInTime, string notes);
    event AttendanceUpdated(uint256 indexed id, uint256 indexed employeeId, uint256 checkInTime, string notes);
    event AuthorizationUpdated(address indexed entity, bool isAuthorized);

    // Modifiers
    modifier onlyOwner() {
        require(msg.sender == owner, "only owner can perform this action");
        _;
    }

    modifier onlyAuthorized() {
        require(authorizedEntities[msg.sender], "not authorized to perform this action");
        _;
    }

    constructor() {
        owner = msg.sender;
        authorizedEntities[owner] = true;
    }

    // Function to authorize an account
    function authorizeAccount(address entity, bool isAuthorized) public onlyOwner {
        authorizedEntities[entity] = isAuthorized;
        emit AuthorizationUpdated(entity, isAuthorized);
    }

    // Function to record attendance
    function recordAttendance(
        uint256 id,
        uint256 employeeId,
        uint256 checkInTime,
        string memory notes
    ) public onlyAuthorized {
        attendanceRecords.push(AttendanceRecord(id, employeeId, checkInTime, notes));
        emit AttendanceRecorded(id, employeeId, checkInTime, notes);
    }

    function updateAttendance(
        uint256 id,
        uint256 employerId,
        uint256 checkInTime,
        string memory notes
    ) public onlyOwner {
        for (uint256 i = 0; i < attendanceRecords.length; i++) {
            if (attendanceRecords[i].id == id) {
                attendanceRecords[i] = AttendanceRecord(id, employerId, checkInTime, notes);
                emit AttendanceUpdated(id, employerId, checkInTime, notes);
                return;
            }
        }

        revert("record not found");
    }

    // Function to get attendance records by given input
    function getAttendance(
        uint256 employerId,
        uint256 fromDate,
        uint256 toDate
    ) public view returns (AttendanceRecord[] memory) {
        AttendanceRecord[] memory tempRecords = new AttendanceRecord[](attendanceRecords.length);
        uint256 count = 0;

        for (uint256 i = 0; i < attendanceRecords.length; i++) {
            bool matches = true;

            if (employerId != 0 && attendanceRecords[i].employerId != employerId) {
                matches = false;
            }

            if (fromDate > 0 && attendanceRecords[i].checkInTime < fromDate) {
                matches = false;
            }

            if (toDate > 0 && attendanceRecords[i].checkInTime > toDate) {
                matches = false;
            }

            if (matches) {
                tempRecords[count] = attendanceRecords[i];
                count++;
            }
        }

        AttendanceRecord[] memory filteredRecords = new AttendanceRecord[](count);
        for (uint256 i = 0; i < count; i++) {
            filteredRecords[i] = tempRecords[i];
        }

        return filteredRecords;
    }
}
